# Start a new ssh-agent

Sometimes the `ssh-agent` stops working or never starts. This can happen for
a variety of reasons. When it does, the fix is simple:

```bash
exec "$(ssh-agent)"
```

Then you can `ssh-add` successfully!

    #bash #ssh #tips
