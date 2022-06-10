# Find files between dates

> Applies to GNU `find` 4.3.3 and up

To find files (and directories) changed between `1 Aug 2022` and `1 Sep 2022`:
```bash
find -newerct "1 Aug 2022" ! -newerct "1 Sep 2022"
```

The `c` in `-newerct` could be replaced with:
```
a   The access time of the file reference
B   The birth time of the file reference
c   The inode status change time of reference
m   The modification time of the file reference
t   reference is interpreted directly as a time
```

This is documented in the man pages as `newerXY`

    #gnu #find #tips
