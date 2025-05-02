# Flags after positional arguments in bash and getopts

Elevator pitch: you have a script which accepts certain options/flags, but you
want them to also work when weaved in between positional arguments.

Using `getopts` alone will not accomplish this, since it will stop processing
flags the moment it encounters a positional argument, as described in POSIX
Chapter 12.2 Guideline. For example:

```sh
#!/bin/sh

set -- -a -b pos1 -c

out=''
while getopts 'abc' opt; do
  case "$opt" in
  a|b|c)
    out="${out}${opt}"
    ;;
  *) ;;
  esac
done

# Will output "ab"
echo "$out"
```

The result we actually want is for `-c` to also be recognized as an option.
Imagine something like this, which all works:

```sh
kubectl get -n mynamespace node mynode
kubectl get nodes -n mynamespace mynode
kubectl get nodes mynode -n mynamespace
```

In fact, the inspiration for this came when I was trying to write a simple
`kubectl` plugin which allows for this, since it's so common to provide the
`-n` argument at different places at different times (expecting the same
result).

Now if you're like me (someone who can't resist a "useful exercise"), you
immediately think "what if I just rearranged the options so the flags were
first? That would be so easy!" And of course, it isn't. And of course, there is
a much better way.

For all the code and tests, see the [source](./src).

## The short answer

Write your own flag parser. IT IS NOT THAT BAD! You can even give yourself
wonderful features like "long options" and optional values for options (if you
choose to ignore Guideline 7--don't). The easiest way is to sacrifice Guideline
5, which says that multiple short options can be provided with one hyphen (e.g.
`-abc`), which is increasingly ignored anyway by programs in the wild which
support `-longflag` options. Go programs are notorious for this. A parser which
does not support multiple short options under one hyphen is actually very
straightforward in `bash`:

```bash
declare -a args
while [[ "$#" -gt 0 ]]; do
  case "$1" in
  -n | --namespace)
    # Requires a value
    namespace=$2
    shift
    ;;
  -v | --verbose)
    # Boolean/switch
    verbose=true
    ;;
  --)
    # Guideline 10, all args after -- are positional
    # Use shift here first if you don't want the -- itself included
    args+=("$@")
    break
    ;;
  -*)
    echo "unknown flag: $1"
    exit 1
    ;;
  *)
    args+=("$1")
    ;;
  esac
  shift
done
set -- "${args[@]}"

# At this point, "$@" are the positional arguments
```

Altogether, the relevant code is only ~10 lines longer than using `getopts`
(the `--` case and `declare`/`set`/`shift`), but you get short and long options
that work anywhere in the argument list.

If you choose to support Guideline 5, it's only ~12 more lines:

```bash
declare -a args
while [[ "$#" -gt 0 ]]; do
  # Guideline 5: multiple short options after a hyphen: -abc
  declare -a parts
  if [[ "$1" =~ ^-[^-] ]]; then
    for ((i = 1; i < ${#1}; i++)); do
      parts+=("-${1:$i:1}")
    done
  else
    parts=("$1")
  fi

  for part in "${parts[@]}"; do
    case "$part" in
    -n | --namespace)
      # Requires a value
      namespace=$2
      shift
      break
      ;;
    -v | --verbose)
      # Boolean/switch
      verbose=true
      ;;
    --)
      # Guideline 10: all args after -- are positional
      # Use shift here first if you don't want the -- itself included
      args+=("$@")
      break 2
      ;;
    -*)
      echo "unknown flag: $1"
      exit 1
      ;;
    *)
      args+=("$1")
      ;;
    esac
  done
  shift
done
set -- "${args[@]}"
```

## A useful exercise

One way to accomplish this while still leveraging `getopts` is to simply
rearrange the arguments so that the options (and their values, if applicable)
will be presented before the positional arguments. For the example above, this
would mean rearranging the arguments as follows:

```txt
# Before
-a -b pos1 -c

# After
-a -b -c pos1
```

We'll write a function for this: `optionsfirst`. However, `$@` is scoped to the
function call, and calling `set` in the function will not affect the `$@` used
by the program (caller) itself. So we need the function to emit the rearranged
arguments, and then we call `set` with the result. Importantly, we need to
ensure that word splitting is not impacted. One might consider using a NUL
delimited string, like this:

```sh
optionsfirst() {
  # generate $newargs array
  # ...
  printf "%s\0" "${newargs[@]}"
}
```

However, there are huge problems with this approach:

1. You cannot call shell builtins with `xargs`, including `set`. Therefore, you
   need to call `set -- "${arr[@]}"` instead.

    ```sh
    $ optionsfirst | xargs -0 set --
    xargs: set: No such file or directory
    ```

2. Command substitution will strip out the NUL characters, so when you try to
   assign the NUL delimited output to a variable, the characters will be
   stripped. You can get around this by piping the output (or using `< <(cmd)`
   syntax) to `readarray`, however, that starts to get needlessly complex for
   the caller. See the examples below on how the NUL stripping works:

    ```bash
    $ a=$(printf "%s\0" arg1 arg2)
    bash: warning: command substitution: ignored null byte in input
    $ xargs -n1 -0 echo <<<"$(printf "%s\0" arg1 arg2)"
    bash: warning: command substitution: ignored null byte in input
    arg1arg2

    $ xargs -n1 -0 echo < <(printf "%s\0" arg1 arg2)
    arg1
    arg2
    $ printf "%s\0" arg1 arg2 | xargs -n1 -0 echo
    arg1
    arg2
    ```

2. Less shocking but much more important: this will not work if one of the
   arguments to the program *should be or should contain* the NUL character,
   which might be more common than you think! The same problem occurs for other
   possible delimiters, even more commonly (think `\n`).

The only truly reliable answer is to set some array available outside the
function, then use `set` with the array. Arrays are not POSIX, so I'm jumping
straight to requiring `bash`:

```bash
optionsfirst() {
  # generate $newargs array
  # ...
  args=("${newargs[@]}")
}

declare -a args
optionsfirst
set -- "${args[@]}"
```

But since we have the features of `bash`, we can offer more flexibility (and
kind of mimic the `getopts` syntax) by letting the user choose a name for the
array, and then assign the variable based on the name they provide. Like this:

```bash
optionsfirst() {
  # generate $newargs array
  # ...

  # Setting "out" will actually set the variable with the name provided in $1
  declare -n out=$1
  out=("${newargs[@]}")
}

declare -a args
optionsfirst args
set -- "${args[@]}"
```

One more implementation detail: how do we know whether an option requires a
value (e.g. an `OPTARG`)? This is important, because when we rearrange the
arguments, we need to make sure to move the values with the flag, if present.
Don't try to guess, there are too many edge cases. Instead, let the user define
them.


```bash
# Options a and b require values, the rest do not
optionsfirst 'ab' args
```

This works (just check if the character is present), but since I wanted to
support `--long-options` as well, I decided to instead use a colon-delimited
string of flags (with leading hyphens):

```bash
optionsfirst '-a:-b:--long-option'
```

Now for the important part--implementation!

```bash
#!/bin/bash

# usage: optionsfirst OPTSTR ARRAYNAME ARGS...
#
# optionsfirst sets the values in the array variable named ARRAYNAME to the
# provided ARGS, starting with all options (e.g. flags) and their values (if
# applicable) first, followed by the positional arguments. OPTSTR is a
# colon-delimited list of flags which require a value (e.g. those that would
# have an OPTARG in getopts). ARRAYNAME should be declared before this function
# is called. If an argumenent with the value "--" is encountered, it and every
# argument after it will be presented in the original order. If the final
# argument expects a value, an explicit empty string argument will be added to
# ensure a positional parameter isn't mistaken for a value.
function optionsfirst() {
  local optstr=$1
  declare -n arrayname=$2
  shift 2
  declare -a options
  declare -a positional
  local onval
  local allpos
  for arg in "$@"; do
    # Always positional after --
    if [[ "$allpos" == yes ]]; then
      positional+=("$arg")
      continue
    fi

    # Value for flag
    if [[ "$onval" == yes ]]; then
      options+=("$arg")
      unset onval
      continue
    fi

    # On separator
    if [[ "$arg" == '--' ]]; then
      positional+=("$arg")
      allpos=yes
      continue
    fi

    # Positional
    if ! [[ "$arg" =~ ^- ]]; then
      positional+=("$arg")
      continue
    fi

    # Flag
    options+=("$arg")
    # shellcheck disable=SC2076
    if [[ ":$optstr:" =~ ":$arg:" ]]; then
      onval=yes
    fi
  done

  # Last arg expects value
  if [[ "$onval" == yes ]]; then
    options+=('')
  fi

  # shellcheck disable=SC2034
  arrayname=("${options[@]}" "${positional[@]}")
}
```

Here's the problem with this: first, it's huge. You still need to loop through
the arguments and use `case` statements to parse the args anyways! So this adds
like 50 lines of code instead of the ~25 lines of just writing your own parser.
Second, it is needlessly complex. What started as a simple idea became unwieldy
quickly, and no one will understand it without a docstring. Additionally, the
entire idea itself (rearranging CLI args and hoping you do it all correctly
with no edge cases) is a failing game. It's much easier to write the parser
yourself, since those edge cases are easier to handle directly!

This was a fun problem-solving session, and I learned some things I didn't
know, but ultimately it's a waste of code. Here's the parser I ended up writing
for my application (after many hours researching and trying these things out):

```bash
declare -a args
while [[ "$#" -gt 0 ]]; do
  case "$1" in
  -n | --namespace)
    namespace=$2
    shift
    ;;
  -*)
    echo "error: unknown flag $1" >&2
    exit 1
    ;;
  *)
    args+=("$1")
    ;;
  esac
  shift
done
set -- "${args[@]}"
```

Less than 20 lines. The additional features I wanted to support in my
generalized solution were YAGNI for my specific case anyway. Whatever :)

    #bash #flags #getopts #getopt #tips
