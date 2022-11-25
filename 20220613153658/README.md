# Better git diff highlighting

To get Github style hightlighting of differences (specifically changed
words/phrases being highlighted or "word by word" differences), first
find the `diff-highlight` binary that is included with git:
```bash
find /usr -type f -name diff-highlight
```

Then take that binary and add it to the `.gitconfig` either at the
project or global level:
```conf
[pager]
  diff = /usr/share/git/diff-highlight/diff-highlight | less
```

If the `diff-highlight` script is not executable, use perl:
```conf
[pager]
  diff = perl /usr/share/git/diff-highlight/diff-highlight | less
```

> NOTE: Along with the `diff` option, the `log` and `show` pagers can be
> set here as well.

This does make `git diff` a little slower!

    #git #diff #tips
