# Use :make in vim

Instead of creating functions for linting, formatting, compiling, etc, set up `:make`!

For instance, `shfmt`:
```vim
setlocal makeprg=shfmt\ -w\ %
setlocal errorformat^=%f:\ %o:%l:%c:\ %m
```

Using `:make` has a ton of benefits, but the best is probably that it populates the quickfix list.

For using `:make` in the background, use the `vim-dispatch` plugin's `m<CR>` shortcut.

## Related

* https://gist.github.com/romainl/ce55ce6fdc1659c5fbc0f4224fd6ad29
* `:help quickfix`
* `:help :make`

    #vim #tips #scripts
