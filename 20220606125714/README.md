# Autocmd for all shell files in vim

Sometimes shell files don't have an extension, so we need a different way to determine whether
we are working on a shell file.

This will run `OnOpenBash()` any time a file is opened which has an appropriate shebang as its
first line:
```vim
function OnOpen()
  if getline(1) == "#!/bin/bash"
    " Do stuff for bash files
    call OnOpenBash()
  endif
endfunction

autocmd BufRead,BufNewFile * call OnOpen()
```

This strategy can be used generally to run an `autocmd` based on a file's contents rather than
its extension.

> NOTE: `OnOpen()` will run *every time* any file is opened!

    #vim #tips #autocmd #scripts
