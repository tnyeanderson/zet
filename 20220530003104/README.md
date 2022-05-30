# Custom vim commands with a bang (!)

Some vim commands can be run with or without a bang (`!`):
```vim
:q
:q!
```

To create a custom command that can be run with or without a bang:
```vim
command -bang CommandName action
```

> Remember, user-defined command names must be capitalized

The `<bang>0` keyword expands to `0` or `1` depending on whether the command was called with a bang:
```vim
command -bang HasBang :echo <bang>0

:HasBang
0
:HasBang!
1
```

To pass this to a function:
```vim
function! CalledWithBang(bang)
  if a:bang == 0
    echo "Not called with a bang"
  endif
  if a:bang == 1
    echo "Called with a bang"
  endif
endfunction
command -bang HasBang :execute "call CalledWithBang(<bang>0)"
```

## Related

* `:help <bang>`

    #vim #tips #scripts
