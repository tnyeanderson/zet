# Force time sync on Kindle without internet

When a Kindle is reset or loses power, its internal clock will likely become
out of sync. On a device connected to a network, this is likely to cause
strange issues. On a device which is always in airplane mode, it will cause
annoyances like "sort by recent" not working.

For example, the clock on my Kindle Paperwhite 2 (PW2) with firmware version
5.6.1.0.6 was over two years behind. This meant anytime I would open a book, it
would set the "you opened this" time to what it *thought* was now, but was
recorded as a date two years ago. Of course, books which were added in the last
two years would then always appear at the top of the list, even if other books
had been more recently read.

To fix this, we have to set the correct date and time, which in normal-world
would mean going into settings, clicking set date and time, and moving on with
your life. But that would make too much sense for Amazon, since this device and
firmware version only has the ability to set the time, not the date (I believe
this is true for most/all Kindles but can't confirm).

The standard process for getting the time to sync is "connect it to WiFi and
let it talk to Amazon for like a day, then it should be fine"... however this
is not an acceptable answer to me, as I don't ever want my Kindle talking to
Amazon for any reason. Therefore, I needed to figure out what the Kindle wanted
from Amazon, and provide those things myself to make it work.

## How the Kindle synchronizes time

Here's what the Kindle does:

- When it connects to WiFi, it will immediately try to contact
  `spectrum.s3.amazonaws.com` using HTTP. If it cannot reach this, it will fail
  and say "you connected to WiFi, but internet wasn't available". In Amazon's
  infinite wisdom, it assumes you also want to disconnect from the network!

- Then, it will reach out to `kindle-time.amazon.com` using HTTP, and
  apparently uses the `Date` header in the response to set the time. If the
  time is too far off, NTP won't work, so we'll need to fake this too.

- Then, it will reach out to `ntp-g7g.amazon.com` using NTP for the finer
  adjustments.

## How to fake it

Given the above, to make this happen without the Kindle ever touching the
internet, you'll need:

- A segmented/separate WiFi network with no internet access (in my case, a
  VLAN)

- A DNS server reachable from this network, with which you can add rewrites for
  some domains (in my case, AdguardHome)

- A DHCP server within this WiFi network, configured to provide the DNS server
  above (in my case, using OPNSense)

- A web server you can use to fake responses for the Kindle's requests (such as
  Caddy)

- An NTP server which is reachable from the segmented network (such as `ntpd`)

First, make sure you set up any firewall rules needed so that the Kindle, once
on the network, can talk to the web server on port 80, the DNS server on port
53, and the NTP server on port 123. You may also need to have more permissive
rules for traffic from the NTP server, depending on your situation.

Spin up a webserver to respond to all requests with 200 OK. For example, this
could be a Caddy server with the following config:

```txt
spectrum.s3.amazonaws.com:80 {
  respond 200
}

kindle-time.amazon.com:80 {
  respond 200
}
```

Set up DNS rules for `spectrum.s3.amazonaws.com` and `kindle-time.amazon.com`
to point to your web server's IP. Then set up one more to point
`ntp-g7g.amazon.com` to your NTP server.

Now turn off airplane mode and connect to the segmented WiFi network. You may
get a prompt saying you need to provide further authentication (like a captive
portal), and you have to click "OK" (if you do not, it will disconnect you from
WiFi). When the portal page doesn't load (because there never was one), you can
just hit the home button to leave the browser, and you should stay connected to
WiFi.

If you want, set up a TCP dump on the router to watch the requests come in:

```bash
tcpdump -i vlan1 -vv -X -s0 dst port 123 or port 80
```

The sync should work on the first try, but you might want to toggle airplane
mode one more time and make sure that the subsequent NTP requests have a
reasonable time (in other words, that the `kindle-time` adjustment worked).

At this point I rebooted my device to be safe, but I'm not sure it's necessary.
Now my sorting works properly!

## References

This took a TON of digging to find out, and the following pages were very
helpful:

- Finding older firmware: <https://www.mobileread.com/forums/showthread.php?t=351683>
- Setting date thread: <https://www.mobileread.com/forums/showthread.php?t=325434>

    #kindle #amazon #time #ntp
