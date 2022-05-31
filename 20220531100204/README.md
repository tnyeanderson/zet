# Movements based on a variable in vim

To move right 4 characters in normal mode:
```vim
4l
```

Therefore, in a vim script:
```vim
normal 4l
```

To move to the right `x` number of characters:
```vim
let x = 4
execute "normal" . x . "l"
```

> Remember, dots are used for concatentation in vim

The `execute` statement executes the parsed string, which will be `normal 4l`.

    #vim #scripts #tips
