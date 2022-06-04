# Open external command result in new vim window

A command's output (including `stdout` and `stderr`) can be opened in a small window at the bottom of the screen:
```vim
botright 10new | read !ls -al
```

Explanation:
* `botright` - open the new window on the bottom of the screen, full width
*  `10new` - create a new horizontally split window that is 10 lines tall
* `|` - like a semicolon, just run the next command after the first
* `read` - insert the given file under the cursor in the new window
* `!ls -al` - the command output which is passed as a file to `read`

A longer version of this command could be used in a function, and with more features!
```vim
function! RunTest()
  " Save the output to a variable
  silent! let output = system('runtestcommand')
  " Don't show the command output if it was successful
  if v:shell_error == 0
    echo "test was successful"
  else
    " Only show output for non-zero exit codes
    " Open a new full-width window on the bottom
    silent! botright 10new
    " Put the contents of the output variable into the new window buffer
    silent! put =output
    " Set options to make the window a scratch space
    setlocal nobuflisted buftype=nofile bufhidden=wipe noswapfile nomodifiable
    " Set the file name (scratch title)
    execute("file test exited with error code: " . v:shell_error)
  endif
endfunction
```

This is kind of like the help or quick fix window, but with whatever you want!

    #vim #scripts #tips
