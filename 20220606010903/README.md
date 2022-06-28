# Text wrapping in vim

To wrap text at 72 characters when typing in insert mode:
```vim
set textwidth=72
```

To re-wrap text that has already been typed, visually select it and type `gw`.
So to do the whole document:
```vim
ggvGgw
```

Or use motion:
```vim
gggwG
```

> NOTE: `gq` and `gw` both do the same thing and and can act visually or with
> motion, but `gw` keeps the cursor where it started and does not use `formatprg`
> or `formatexpr`

The text will be wrapped either at `textwidth` or the width of the current
window (up to 80 columns).

    #vim #tips #formatting
