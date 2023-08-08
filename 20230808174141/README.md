# Create test files of arbitrary size with dd

The `dd` utility is obviously useful for many things. Sometimes you just need
to try writing a file of a certain size, either to see how long it takes or see
if it's possible on the target filesystem. This is pretty simple:

```bash
# Create a 1MB file "output.txt" and fill it with zeros.
dd if=/dev/zero of=output.txt bs=1k count=1024
```

- The `bs` option is "byte size", and `count` is the number of these units you
  want to create. So the total size of the file will be `count * bs`
- The default byte size of 512 can be pretty slow. Try increasing it.
- Try using `/dev/random` as the input if your are testing a filesystem which
  uses compression.
- Add the `--progress` flag for a progress indicator.
- Run this command through the `time` utility to see how long it takes to write
  the file.

    #bash #dd #linux #fs
