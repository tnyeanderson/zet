# Force exit from a hung SSH session

The secret is to type the following in order:

`Enter` -> `~` -> `.`

> NOTE: Add an additional tilde for each nested SSH session

This is known as an "escape sequence" and these can be enumerated with:

`Enter` -> `~` -> `?`

Example:
```
 ~.   - terminate connection (and any multiplexed sessions)
 ~B   - send a BREAK to the remote system
 ~C   - open a command line
 ~R   - request rekey
 ~V/v - decrease/increase verbosity (LogLevel)
 ~^Z  - suspend ssh
 ~#   - list forwarded connections
 ~&   - background ssh (when waiting for connections to terminate)
 ~?   - this message
 ~~   - send the escape character by typing it twice
(Note that escapes are only recognized immediately after newline.)
```

Optionally, set `ServerAliveInterval 5` in `.ssh/config` to automatically end
sessions after 5 seconds of no data received.

## Related

- `man ssh` -> `/ESCAPE CHARACTERS`

    #ssh #tips
