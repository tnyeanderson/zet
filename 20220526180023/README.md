# Searching for functions in vim :help

Use parenthesis to denote a function when searching vim `:help`.
```vim
:help match()
```

The following *wrong way* will return `:match`, which is *not* the same as the `match()` function:
```vim
" This will not return what we are looking for
:help match
```

To search all functions:
```vim
:help function()

" Refine results by searching. Lines begin with function names.
/^match
```

    #vim #tips #rtfm
