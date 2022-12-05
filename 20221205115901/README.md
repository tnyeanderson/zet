# Completion for docker compose alias

Since `docker compose` is now a subcommand of `docker`, getting bash completion
working is a little painful.

The reason for this pain is that the upstream completion function is expecting
the base command to be `docker`, and `compose` to be one of its subcommands
(two words). But since the alias is only one word, it screws up.

The answer is to create and assign a new completion for the alias:
`c`):

```bash
#!/bin/bash

# set the alias
alias c="docker compose"

# initial setup, make sure you do this somewhere
source /etc/profile.d/bash_completion.sh

# since c is aliased to a subcommand, fake the upstream completion
_docker_compose_alias() {
	local result=$(command "$COMPOSE_PLUGIN_PATH" __completeNoDesc compose "${COMP_WORDS[@]:1}" 2>/dev/null | grep -v '^:[0-9]*$');
	COMPREPLY=($(compgen -W "${result}" -- "$current"))
}

# enable the completion
complete -F _docker_compose_alias c
```

This was derived from the upstream `_docker_compose` completion function:

```bash
_docker_compose ()
{
    local completionCommand="__completeNoDesc";
    local resultArray=($COMPOSE_PLUGIN_PATH $completionCommand compose);
    for value in "${words[@]:2}";
    do
        if [ -z "$value" ]; then
            resultArray+=("''");
        else
            resultArray+=("$value");
        fi;
    done;
    local result=$(eval "${resultArray[*]}" 2> /dev/null | grep -v '^:[0-9]*$');
    COMPREPLY=($(compgen -W "${result}" -- "$current"))
}
```

    #bash #completion #docker #compose #alias
