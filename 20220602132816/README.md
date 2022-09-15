# Repeatable, idempotent ln

Say you create a soft link:
```bash
ln -s src target
```

If you try to repeat that same action, it will fail:
```bash
ln -s src target
# ln: failed to create symbolic link 'target': File exists
```

Adding `-f` will force the link to be overwritten:
```bash
ln -sf src target
```

This works for files, but what if `src` is a directory?
```bash
ln -sf src target
# .
# ├── src
# └── target -> src
ln -sf src target
# .
# ├── src
# │   └── src -> src
# └── target -> src
```

That's not correct! Since `target` resolves to a directory (`src`), the link gets created *inside* that directory.

This can be avoided with the `-n` flag:
```
-n, --no-dereference
treat LINK_NAME as a normal file if it is a symbolic link to a directory
```

> NOTE: While the `-T` flag accomplishes a similar goal, it is not available on all systems. `-n` is most compatible.

Let's try again:
```bash
ln -sfn src target
# .
# ├── src
# └── target -> src
ln -sfn src target
# .
# ├── src
# └── target -> src
```

Therefore, for a truly idempotent and repeatable `ln` command, always use:
```bash
ln -sfn src target
```

> NOTE: If the user has an existing file or link in the target location, it *will* be overwritten with this strategy!

    #bash #tips #scripts
