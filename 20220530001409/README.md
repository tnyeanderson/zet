# Echo current filename in vim

If using an external command (`:!`), the `%` keyword can be used directly:
```vim
:!echo %
```

If using a vim command, interpolate with `expand(%)`:
```vim
:echo "Current filename: " . expand('%')
```

Or use the special `%` register using the `@` accessor:
```vim
:echo "Current filename: " . @%
```

## Related

* `:help cdline-special`

    #vim #tips #scripts
