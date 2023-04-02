# Configure Dell BIOS from linux

Dell has an impressively-poorly-named utility called `Dell Command | Configure`
that allows managing BIOS configurations from the OS. However, Dell makes it
annoyingly difficult to find the download links for the different releases and
platforms.

After searching forever, I finally found a [random stackoverflow
answer](https://superuser.com/a/1648187) which has the KB article which
*actually lists the download links*. Most others force you to "choose a
compatible device", but of course I don't know which device to pick to get the
latest Ubuntu release of the software, since my hardware is rather old and only
gives the Windows versions (Optiplex 7040).

Below are the links for version `4.10.0`, released January 2023:

- [Ubuntu Desktop 18.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=2M1YY)
- [Ubuntu Server 18.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=2M1YY)
- [Ubuntu Desktop 20.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=TX7GF)
- [Ubuntu Desktop 22.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=8VCG0)
- [RHEL 7.x (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=J4GY5)
- [RHEL 8.0 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=7FXYJ)

Because I don't trust Dell to keep these available, I've downloaded them all
and put them in the `downloads/4.10.0` directory of this zet.

> NOTE: I do not own the copyright to these files, I am only providing a
personal backup copy of the software Dell has publicly distributed.

To use it on Ubuntu Server 18.04 (which is recommended as the old libssl
dependency is not available on Ubuntu 22):

1. Download the `.tar.gz` file from either the KB article, links above, or the
   `downloads` directory of this zet
2. Extract it
3. Install the `srvadmin-hapi` deb package with `dpkg -i srvadmin-hapi*.deb`
4. Install the `command-configure` deb package with `dpkg -i
   command-configure*.deb`
5. Use the tool at `/opt/dell/dcc/cctk`

Find available commands:

```bash
/opt/dell/dcc/cctk -H
```

Export your current configuration:

```bash
/opt/dell/dcc/cctk -O mybios.conf
```

Import a previously exported configuration:

```bash
/opt/dell/dcc/cctk -I mybios.conf
```

I really wish Dell would make this stuff easy to find. I literally spent a week
looking for those download links and had given up, only to stumble the KB
article when looking for other BIOS configuration software!

    #bios #dell #hardtofind #config
