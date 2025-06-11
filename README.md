# :unlock: TrustTap

Send a link to your friend. When they tap it, you whitelist their IP address.

I've been using [Nginx Proxy Manager](https://nginxproxymanager.com/) and [NPMPlus](https://github.com/ZoeyVid/NPMplus) for a bit, big fan of their Access List feature.
But handling the whitelisting of IP addresses is a bit of a pain, so I wrote this. It's a bunch of Go, probably a little
over-engineered, but it works.

## Building

```shell
podman build -f build/package/Dockerfile .
```
