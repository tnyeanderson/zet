# Exit code of external command in vim

To get the exit code of an external command, use `v:shell_error`:
```vim
execute("!which git")
if v:shell_error == 0
  echo "git command exists"
else
  echo "git command does not exist. exit code: " . v:shell_error
fi
```

    #vim #tips #scripts
