# Running git restore after git add

Say you have run `git add` on a file, then you changed the file further (or
even deleted it). How does the `git restore` command act in this situation?

For a demonstration, a new repo was created and a `test` file was committed to
it with the following contents:

```
first commit
```

Now, `test` is edited to instead contain `second commit`, and `git add` is run:

```
$ echo 'second commit' >test
$ git add test
$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
  modified:   test
```

The changes are now staged. What if we delete the file?
```
$ rm test
$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
  modified:   test

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
  deleted:    test
```

So now, the modified file is staged, but the file is deleted from the worktree.
To restore the staged (modified) version of the file, use plain old `git
restore`:

```
$ git restore test
$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
  modified:   test
```

> WARNING: If you use `git restore --staged test` instead, your original staged
changes will be gone forever! This removes the changes from the staging area,
and your current worktree has deleted the file, so the changes are now lost!

    #git #tips #undo
