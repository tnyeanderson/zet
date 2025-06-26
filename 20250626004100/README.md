# Controlling or inspecting traffic with mitmproxy

Often, I want to inspect or tightly control network or HTTP traffic. An easy
example is when you are developing some API or software, and you want to see
*exactly* what the client called and how the server responded.

Another example: a printer from a stupid company that doesn't just release
firmware blobs I can provide to the printer, but instead *only* allows the
printer to be updated using a proprietary tool that requires Windows, or by
telling the printer itself to reach out to the internet for an update.
Obviously, *printers shouldn't be reaching out to the internet*, so I've
restricted my printer behind a firewall to prevent this kind of stuff.

The second example is much more complex than the first, but they both have the
same answer: use `mitmproxy`!

<https://mitmproxy.org/>

Start by installing the program, then run it. No arguments are needed, unless
you need to add `-k` to support self-signed certificates upstream like I did,
because companies that make printers are extraordinarily braindead. Before you
jump in to correct me, yes, I understand the reason they did this. And no, I
don't think it's a good reason.

Next, configure your client, in this case the printer, to send its traffic
through the proxy (which is listening on port `8080`). You'll immediately see
the requests start coming in, and you can inspect them thoroughly as they pass
through the proxy!

Of course, be sure to follow the instructions for trusting the `mitmproxy` CA
if your clients are making `https` calls. Usually this can be retrieved with
the following command:

```sh
http_proxy=http://proxyhost:8080 curl http://mitm.it/cert/cer
```

Then add the provided cert as a trusted CA on the client device. **Remember to
un-trust it when you are done!**

To block ALL incoming traffic by default, such that you must manually allow
both the request *and* the response through in real time, use this:

```sh
./mitmproxy --intercept '~all'
```

> NOTE: You can also edit this intercept config by pressing `i` in the
> interactive console.

If you want to block all traffic except for a specific URL, do something like this:

```sh
# Gee I wonder what the stupid company was that made me do this stupid extra work
./mitmproxy --intercept '!(~u cu\.bwc\.brother\.com/certset/ver)'
```

See the [docs](https://docs.mitmproxy.org/stable/concepts/filters/) for more
filter expressions.

When traffic is intercepted, you must select the flow in the console and click
`a` to let the request go through to the upstream server, then `a` again to
allow the response to be sent to the downstream client. This allows you to
easily see what requests will be made *before* you blindly allow them, avoiding
the possibility of inadvertently allowing data exfiltration via a very "smart"
device. It also allows you to scrutinize what the server sends back, in case
that is important.

The program is extremely powerful and this just scratches the surface of its
possibilities. Have fun!

> P.S. - If you are looking for a much worse version of this, check out my
> `flies` project! It's no longer maintained, so you know it's rock solid
> stable! It was fun to make, but `mitmproxy` makes it worthless. Don't use it.
>
> <https://github.com/tnyeanderson/flies>

    #http #proxy #net #iot #security #api
