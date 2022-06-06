# Text wrapping in vim

To wrap text at 72 characters when typing in insert mode:
```vim
set textwidth=72
```

To re-wrap text that has already been typed, visually select it and type `gq`.
So to do the whole document:
```vim
ggvGgq
```

The text will be wrapped either at `textwidth` or the width of the current window (up to 80 columns).

    #vim #tips #formatting
