# Planetmint Faucet

Planetmint faucet is a Go web server for requesting `plmnt` tokens from a dedicated wallet.

## Getting started

Build it:
```
$ go build -v ./cmd/...
```

Run it:
```
$ ./faucet
```

Use it:
- Open web browser at [http://localhost:8080/](http://localhost:8080/).

## Development

Format HTML code, e.g.:
```
$ tidy -i -m -w 160 -ashtml -utf8 templates/index.html
```

## Configuration

The service needs to be configured via the `./app.toml` file or environment variables. The defaults are:
```toml
address = 'plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty'
service-bind = 'localhost'
service-port = 8080
```
