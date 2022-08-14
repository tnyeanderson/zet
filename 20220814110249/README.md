# Wiping disks

To wipe a disk and make it (near) impossible to recover the data, the
long-standing choice is DBAN. The open source alternative is ABAN.

Both of the above tools are bootable images. To wipe a disk in regular old
linux, use `shred` on the entire device:
```bash
# One pass of random data, final pass of zeros
sudo shred --verbose --random-source=/dev/urandom -n1 --zero /dev/sdX
```

> NOTE: `shred` can be used for files or devices, but it only works on certain
(non-CoW) file systems. Read the man pages and use caution!!

## Related

- `man shred`
- https://wiki.archlinux.org/title/Securely_wipe_disk
- https://aban.derobert.net/
- https://dban.org/
