# UniFi Network Application login error

You may see the following message when trying to log in to UniFi Network
Application (the successor to UniFi Controller):

> Login error
> 
> There was an error making that request. Please try again later.

If you are using Caddy v2.11.0 or above, you may need to set this `header_up`
configuration:

```txt
reverse_proxy my-backend {
  header_up Host {hostport}
}
```

I found that I could recreate browser's request in curl and reproduce the 403
response:

```bash
curl -vvv \
  -H 'Origin: https://unifi.example.com' \
  -H <other_header>... \
  --data '{"username":"myuser","password:"mypass","strict":true}' \
  https://unifi.example.com/api/login
```

If I removed the `Origin` header, it returned the expected 200. That led me to
this issue:

<https://github.com/nixos/nixpkgs/issues/498025>

Which links to the upstream documentation explaining:

> Since (most) headers retain their original value when being proxied, it is
> often necessary to override the Host header with the configured upstream
> address when proxying to HTTPS, such that the Host header matches the TLS
> ServerName value ... Since Caddy v2.11.0, this is done automatically, so it is
> no longer necessary to explicitly override the Host header when proxying to
> HTTPS. If you wish to opt out of this behavior, you can set the Host header to
> its original value (but this rarely makes sense to do).

<https://caddyserver.com/docs/caddyfile/directives/reverse_proxy#https>

The UniFi application absolutely sucks when it comes to debugging these issues.
There's zero helpful logs, and the source code isn't available. This was found
through hours of investigation, searching, and trial and error.

    #unifi #caddy #bug #auth
