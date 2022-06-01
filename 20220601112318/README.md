# Use ddrescue for failing drives

If a drive is failing, don't try to use `dd` to image the drive for replacement or backup. If `dd` encounters
errors, it will not copy that data, and could actually make the entire image unusable!

Instead, the `ddrescue` tool should be used. Say there is a failing drive `/dev/sda` that needs to be imaged
to a new/replacement drive `/dev/sdb`:
```bash
ddrescue -f /dev/sda /dev/sdb /root/recovery.log
```

Or to image `/dev/sda` to a file (`/root/drive.img`):
```bash
ddrescue -f /dev/sda /root/drive.img /root/recovery.log
```


Options:
* `-f`: Force overwrite the destination file (required for disks)
* `-n`: Short for `-no-scrape` which skips trying to recreate heavily damaged areas of a file
* `/root/recovery.log`: Path to the log file for the recovery

    #disaster #recovery #tips #disks
