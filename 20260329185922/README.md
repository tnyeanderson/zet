# UniFi Network Application log level

In `/usr/lib/unifi/data/system.properties`, you can manipulate the log levels.
If using `linux-server/unifi-network-application`, this location is equivalent
to `/config/data/system.properties`.

For example:

```txt
debug.mgmt=debug
debug.device=debug
debug.system=debug
debug.sdn=debug
```

The default value for each of these appears to be `warn`. None of these options
are documented by upstream as far as I can tell.

Increasing these log levels does not seem to give any more insight into why
`/api/login` returns 403 for a valid user.

    #unifi #log #homelab
