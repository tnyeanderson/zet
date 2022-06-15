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

For per-process usage (using `tmux` as an example):
```bash
sudo lsof -p $(pidof tmux)
```

Example output:
```
COMMAND     PID   USER   FD   TYPE             DEVICE SIZE/OFF     NODE NAME
tmux:\x20 73312 thomas  cwd    DIR                8,2     4096 11534339 /home/thomas
tmux:\x20 73312 thomas  rtd    DIR                8,2     4096        2 /
tmux:\x20 73312 thomas  txt    REG                8,2   692192  4067537 /usr/bin/tmux
tmux:\x20 73312 thomas  mem    REG                8,2  5699248  4073195 /usr/lib/locale/locale-archive
tmux:\x20 73312 thomas  mem    REG                8,2   157224  4068628 /usr/lib/x86_64-linux-gnu/libpthread-2.31.so
tmux:\x20 73312 thomas  mem    REG                8,2  2029592  4067540 /usr/lib/x86_64-linux-gnu/libc-2.31.so
tmux:\x20 73312 thomas  mem    REG                8,2   101352  4068630 /usr/lib/x86_64-linux-gnu/libresolv-2.31.so
tmux:\x20 73312 thomas  mem    REG                8,2   346672  4071319 /usr/lib/x86_64-linux-gnu/libevent-2.1.so.7.0.0
tmux:\x20 73312 thomas  mem    REG                8,2   192032  4072338 /usr/lib/x86_64-linux-gnu/libtinfo.so.6.2
tmux:\x20 73312 thomas  mem    REG                8,2    14568  4068300 /usr/lib/x86_64-linux-gnu/libutempter.so.1.1.6
tmux:\x20 73312 thomas  mem    REG                8,2    14880  4068652 /usr/lib/x86_64-linux-gnu/libutil-2.31.so
tmux:\x20 73312 thomas  mem    REG                8,2    51856  4068618 /usr/lib/x86_64-linux-gnu/libnss_files-2.31.so
tmux:\x20 73312 thomas  mem    REG                8,2    27002  4599017 /usr/lib/x86_64-linux-gnu/gconv/gconv-modules.cache
tmux:\x20 73312 thomas  mem    REG                8,2   191504  4064382 /usr/lib/x86_64-linux-gnu/ld-2.31.so
tmux:\x20 73312 thomas    0u   CHR                1,3      0t0        6 /dev/null
tmux:\x20 73312 thomas    1u   CHR                1,3      0t0        6 /dev/null
tmux:\x20 73312 thomas    2u   CHR                1,3      0t0        6 /dev/null
tmux:\x20 73312 thomas    3r  FIFO               0,13      0t0   583043 pipe
tmux:\x20 73312 thomas    4w  FIFO               0,13      0t0   583043 pipe
tmux:\x20 73312 thomas    5u   CHR              136,1      0t0        4 /dev/pts/1
tmux:\x20 73312 thomas    6u  unix 0xffffffffffffffff      0t0   583044 /tmp/tmux-1000/default type=STREAM
tmux:\x20 73312 thomas    7u  unix 0xffffffffffffffff      0t0   582078 type=STREAM
tmux:\x20 73312 thomas    9u   CHR                5,2      0t0       90 /dev/ptmx
tmux:\x20 73312 thomas   11u   CHR                5,2      0t0       90 /dev/ptmx
tmux:\x20 73312 thomas   13u   CHR                5,2      0t0       90 /dev/ptmx
```

## Related

- https://www.kernel.org/doc/html/latest/admin-guide/sysctl/fs.html
- https://www.baeldung.com/linux/error-too-many-open-files

    #kernel #linux #ulimit
