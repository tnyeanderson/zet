# Prepend value to :set string in vim

In the same way that `+=` can *append* a string option in vim:
```vim
" Equivalent to var=existing+new
set errorformat+=%f:%l:%m
```

Use `^=` to *prepend* a string option:
```vim
" Equivalent to var=new+existing
set errorformat^=%f:%l:%m
```

>NOTE: In both circumstances, if the option is a comma separated list, the comma will be added automatically.

    #vim #tips #scripts
