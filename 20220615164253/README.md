# File descriptor limits in linux (ulimit)

To see the current system-wide limits:
```bash
cat /proc/sys/fs/file-nr
```

This will give 3 numbers:
```
24832   0       9223372036854775807
```

In order, they represent:
- Number of allocated file handles
- Number of allocated but unused file handles
- Maximum number of file handles

> NOTE: The second number (allocated but unused) will always report `0` in
> Linux 2.6+, which just means that the number of allocated file handles
> exactly matches the number of used file handles.

To see the limits for the current user:
```bash
cat /proc/self/limits
```

Example output:
```
Limit                     Soft Limit           Hard Limit           Units
Max cpu time              unlimited            unlimited            seconds
Max file size             unlimited            unlimited            bytes
Max data size             unlimited            unlimited            bytes
Max stack size            8388608              unlimited            bytes
Max core file size        0                    unlimited            bytes
Max resident set          unlimited            unlimited            bytes
Max processes             127881               127881               processes
Max open files            1024                 1048576              files
Max locked memory         unlimited            unlimited            bytes
Max address space         unlimited            unlimited            bytes
Max file locks            unlimited            unlimited            locks
Max pending signals       127881               127881               signals
Max msgqueue size         819200               819200               bytes
Max nice priority         30                   30
Max realtime priority     99                   99
Max realtime timeout      unlimited            unlimited            us
```

## Related

- https://www.kernel.org/doc/html/latest/admin-guide/sysctl/fs.html

    #kernel #linux #ulimit
