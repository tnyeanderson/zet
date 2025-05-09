# If line matches/contains string in vim script

Say you have a function in your `.vimrc` or some other `.vim` script. In this function you want a conditional
tocheck if a line contains text, and take an action if it does. This is perfect for the `match` function!

```vim
:help match()
:echo match("testing", "ing")   " results in 4
:echo match("testing", "ed")   " results in -1
```

>HINT: Use `getline('.')` to get the text in the current line!

Therefore:
```vim
" If current line contains 'hello'
if match(getline('.'), 'hello') != -1
	doStuff
fi
```

Or use regex:
```vim
" If current line starts with 'hello'
if match(getline('.'), '^hello') != -1
	doStuff
fi
```

Since match returns the index, this could also be written as:
```vim
let line = getline('.')
if match(line, 'hello') == 0
	doStuff
fi
```

    #tips #vim #scripts
