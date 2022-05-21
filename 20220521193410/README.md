# Git diff between repos

To get a diff between two repos (like a hard fork), first clone one of the repos
```bash
git clone https://github.com/user1/repo1.git
```

Then add the other repo as a remote branch:
```bash
git remote add -f user2repo https://github.com/user2/repo2.git
git remote update
```

Then diff the branches:
```bash
git diff origin/master user2repo/master
```

Don't forget to delete the remote when you're done!
```bash
git remote remove user2repo
```

    #git #diff #tips
