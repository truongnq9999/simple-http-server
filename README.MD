# Simple http server

This is a simple http server write in golang. It simply returns request body or request parameters when body is null.

## Usage

```bash
http-server -port 8080 -show-header=true
```

## Options

```bash
  -port int
        Port to listen on (default 8080)
  -show-header
        Show request header
```

## Build

```bash
go build -o http-server
```

## Dockerize

```bash

docker build -t registry.vnpay.vn/vnpay-ocp/efin/share/simple-http-server:1.0.0 .
docker push registry.vnpay.vn/vnpay-ocp/efin/share/simple-http-server:1.0.0
```
