# :unlock: TrustTap

[![Validate](https://github.com/RobKenis/TrustTap/actions/workflows/validate.yaml/badge.svg)](https://github.com/RobKenis/TrustTap/actions/workflows/validate.yaml)

Send a link to your friend. When they tap it, you whitelist their IP address.

I've been using [Nginx Proxy Manager](https://nginxproxymanager.com/) and [NPMPlus](https://github.com/ZoeyVid/NPMplus) for a bit, big fan of their Access List feature.
But handling the whitelisting of IP addresses is a bit of a pain, so I wrote this. It's a bunch of Go, probably a little
over-engineered, but it works.

## Developing

I am using [templ](https://templ.guide/) to manage Go templates. Install it with:

```shell
go get -tool github.com/a-h/templ/cmd/templ@latest
```

Then run `go tool templ generate` to generate the templates.

### Linting

```shell
pre-commit install
pre-commit run --all-files
```

### Running tests

```shell
go test -cover -coverprofile=coverage.out -parallel 10 ./...
go tool cover -html=coverage.out
```

## Building

```shell
podman build -f build/package/Dockerfile .
```
