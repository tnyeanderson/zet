# Resize/shrink logical volume

When using Proxmox, the web interface for growing a disk frequently confuses me
and I put in the *total* size I want the disk to be, rather than the amount by
which the disk should *grow*.

When this mistake is made, correct it by doing the following:

Find the correct lvm to resize using `lvs`, `vgs`, `pvs`, and/or the web
interface.

Shrink the logical volume to the correct size (in this example, 20GB):
```bash
lvreduce -L 20G pve/vm-114-disk-2
```

Edit the config for the VM to correct the size. The file is stored at
`/etc/pve/qemu-server/$id.conf`.

Confirm the correct disk space from the VM. A reboot may be required.
Alternatively, unmount and *carefully* detach the disk from the web interface,
then reattach it by clicking **Edit > Add**. It is so much less painful to just
reboot. The disk can be remounted after this. **IT IS CRITICAL THAT THE
PARTITION IS NOT RESIZED UNTIL THE CORRECT SIZE IS REFLECTED IN
`lsblk`/`fdisk`/`cfdisk`!!**

*Now* resize the partition on the host using `cfdisk /dev/sdb` (or whatever the
disk is).

Finally, use `resize2fs /dev/sdb1` (or whatever the partition is) to resize the
filesystem on the partition.

    #proxmox #lvm #disk
