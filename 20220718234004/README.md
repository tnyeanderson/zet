# Encrypted ZFS filesystems

Encryption in ZFS is done at the file system level (sometimes called
dataset/volume). Therefore it uses the `zfs` command, in contrast to the
`zpool` command which operates at the pool level.

The steps below will create an encrypted ZFS file system called `secrets` in
a pool named `tank`.

First generate a passphrase (max 128 characters) and store it in a file with
restricted permissions. For example: `/etc/zfs/keys/tank-secrets.key`

Then create the encrypted volume:
```bash
zfs create \
  -o encryption=on \
  -o keylocation=file:///etc/zfs/keys/tank-secrets.key \
  -o keyformat=passphrase \
  tank/secrets
```

- `encryption=on` : Use the default encrytion method (`aes-256-gcm`)
- `keylocation=file:///etc/zfs-secrets.key` : Path to the key file (starting with `file://`)
- `keyformat=passphrase` : Must be between 8 and 512 bytes long

When the volume is created, the key will be loaded and the file system will be
mounted. However, by default the file system will not be available at startup.
A systemd unit must be defined to load the encryption keys at boot.

Create `/etc/systemd/system/zfs-load-key.service`:
```bash
[Unit]
Description=Load encryption keys
DefaultDependencies=no
After=zfs-import.target
Before=zfs-mount.service

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/sbin/zfs load-key -a
# Enable the StandardInput directive if keylocation=prompt
#StandardInput=tty-force

[Install]
WantedBy=zfs-mount.service
```

> NOTE: Make sure the path to the `zfs` binary is correct!

Finally, enable the service:
```bash
systemctl enable zfs-load-key.service
```

Test that it works by rebooting!

## Related

- `man zfsprops`
- https://nsg.cc/post/2022/encrypted-zfs/
- https://arstechnica.com/gadgets/2021/06/a-quick-start-guide-to-openzfs-native-encryption/
