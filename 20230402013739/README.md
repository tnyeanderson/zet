# Configure Dell BIOS from linux

Dell has an impressively-poorly-named utility called `Dell Command | Configure`
that allows managing BIOS configurations from the OS. However, Dell makes it
annoyingly difficult to find the download links for the different releases and
platforms.

After searching forever, I finally found a [random stackoverflow
answer](https://superuser.com/a/1648187) which links to the [KB
article](https://www.dell.com/support/kbdoc/en-us/000178000/dell-command-configure)
which *actually lists the download links*. Most others force you to "choose
a compatible device", but of course I don't know which device to pick to get
the latest Ubuntu release of the software, since my hardware is rather old and
only gives the Windows versions (Optiplex 7040).

Below are the links for version `4.10.0`, released January 2023:

- [Ubuntu Desktop 18.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=2M1YY)
- [Ubuntu Server 18.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=2M1YY)
- [Ubuntu Desktop 20.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=TX7GF)
- [Ubuntu Desktop 22.04 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=8VCG0)
- [RHEL 7.x (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=J4GY5)
- [RHEL 8.0 (64-bit)](https://www.dell.com/support/home/en-us/drivers/DriversDetails?driverId=7FXYJ)

Because I don't trust Dell to keep these available, I've downloaded them all
and put them in the `downloads/4.10.0` and `downloads/4.8.0` directories of
this zet.

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

Ubuntu 22 cannot use version 4.10.0 to a libssl1.1 dependency. Version 4.8.0
embeds its libssl dependency in the /opt/dell folder, so it does work on the
latest Ubuntu 22. However, it can also [screw
up](https://www.dell.com/community/Linux-Developer-Systems/Dell-Command-Configure-breaks-OpenSSL-and-other-things/m-p/8275666)
other parts of the system.

[This forum post](https://bbs.archlinux.org/viewtopic.php?id=280992) explains
the remedy. In summary, to use `cctk` on Ubuntu 22:

1. Install the Ubuntu 20 version of the 4.8.0 package using the same method as
   above
2. Get rid of the extra config created by Dell
   ```bash
   mv /etc/ld.so.conf.d/hapiintfdcc.conf /etc/ld.so.conf.d/hapiintfdcc.conf.bak
   ```
3. Reconfigure shared library paths
   ```bash
   ldconfig
   ```

I really wish Dell would make this stuff easy to find. I literally spent a week
looking for those download links and had given up, only to stumble the KB
article when looking for other BIOS configuration software!

    #bios #dell #hardtofind #config
