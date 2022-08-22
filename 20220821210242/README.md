# Fix intel wifi cards on linux

Trouble connecting to (or staying connected to) wifi networks on linux? I've
had issues with the Intel Centrino Wireless-N 1000 card on Kubuntu 22. The
issue seems to be that the card doesn't negotiate "N" standard wireless
connections well. The fix is to disable the "N" (b/g/n) standard and allow "B"
and "G" (b/g) only.

To disable "N" standard wireless connections, add the following line to
`/etc/modprobe.d/iwlwifi.conf`:
```
options iwlwifi 11n_disable=1
```

Then restart to apply the changes (or remove and re-add the module).

    #tips #linux #wifi #intel
