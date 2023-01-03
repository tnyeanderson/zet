# Fix e1000 network hangs

Under some conditions, intel e1000 and e1000e network cards can cause hangs. To
fix this on debian-ish systems, disable `tso` and max out the `rx` and `tx`
values.

> IMPORTANT: Please ensure you are on an up-to-date kernel before trying this!
Confirmed working with 5.15, and if memory serves 5.12 did NOT work!!

This can be done by creating `/etc/network/if-up.d/e1000e-fix` with the
following contents and making it executable (change the interface name as
needed):
```bash
#!/bin/sh

ethtool -G enp0s31f6 rx 4096 tx 4096
ethtool -K enp0s31f6 tso off
```

This will set the correct parameters once the `enp0s31f6` interface is up.
**Remember to adjust the interface name in the above script as needed!**

    #intel #network #bug #kernel
