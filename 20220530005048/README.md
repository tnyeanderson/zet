# Write file with sudo in vim

Sometimes a file owned by root is opened in vim by a non-root user. If the user tries to `:w` the changes,
vim will refuse. One way around this is to use the `write` command in a special way:
```vim
:w !sudo tee >/dev/null %
```

Explanation:
* `:w` - the `write` command. When called with `:w !cmd`, `cmd` is executed with the contents of the current
  file passed to stdin. In this case, `write` is not acually writing the file, it's passing it to `tee`
  which does the actual writing
* `!sudo tee` - execute the `tee` command with `sudo`
  * `tee` outputs stdin to stdout and writes it to the file given as an argument
* `>/dev/null` - don't show the stdout output of `tee`
* `%` - vim expands this to the current file name or relative path. This is the filename argument to the
  `tee` command

A nice helper function mapped to the command `:W!`:
```vim
" save files using sudo with :W!
function! WriteSudo(bang)
  if a:bang != 1
    echo "Use :W! to save with sudo"
    return
  endif
  silent write !sudo tee >/dev/null %
  silent edit!
  redraw!
  echo @% . " written with sudo"
endfunction
command -bang W :execute "call WriteSudo(<bang>0)"
```

Why this is better than all the other proposed ways:
* Some people recommend using `cmap` (in some capacity) but this is frustrating as it adds delay when typing
  `:w`, replaces the text in the `:` line while typing, then visually prints the expanded command to be run.
* If the user tries to use `:W` without the bang, it tells the user to run `:W!` instead. This forces the user
  to think about it and avoids mistyping `:W` instead of `:w`
* The file is automatically reloaded from disk (avoiding "File has been changed outside vim" errors)
* The file/path is outputted to vim and the user is informed that the file was written with sudo. No other
  proposed method does this (to my knowledge)

   #vim #tips #scripts
