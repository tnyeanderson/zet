# Git commit hooks for a repo

There is no way to *automatically* set up git hooks for anyone who pulls and commits to a repo. This is
done for security reasons (hooks execute arbitrary scripts as a side effect during normal `git` operations).

Therefore, the user must take action to set up the hooks themselves. One way to do this is to ask the user
to run a setup script. For instance, place properly named hook scripts in `.githooks`:
```
repo/
└── .githooks
    ├── pre-commit
    └── setup-hooks.sh
```

`setup-hooks.sh` sets up the `hooksPath`:
```bash
#!/bin/bash

# Set the hooks folder
git config core.hooksPath .githooks
```

Then, in the README, tell the user to run `.githooks/setup-hooks.sh` before contributing to the project.

    #git #tips #cicd #hooks
