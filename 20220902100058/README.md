# View logs for MacOS software update

To see the logs for a failed (or successful) software update for MacOS:
```bash
log show --debug --predicate 'subsystem == "com.apple.SoftwareUpdate"' -last 10m
```

    #apple #macos #debug #logs
