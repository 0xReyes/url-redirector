# url-redirector

A simple Go-based HTTP redirector driven by YAML config.

## Features
- Host-only and host+path redirects  
- Configurable HTTP status codes  
- Zero external dependencies

## Prerequisites
- Go 1.18+  

## Installation
```bash
git clone https://github.com/0xReyes/url-redirector.git
cd url-redirector
go build -o url-redirector main.go
```

## Configuration
Edit **config.yaml**:
```yaml
server:
  address: "0.0.0.0"
  port: 80

redirects:
  "old.example.com":
    to: "https://new.example.com/"
    status: 301
```

## Usage
```bash
./url-redirector
```

## Docker (optional)
```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o url-redirector main.go

FROM alpine
COPY --from=builder /app/url-redirector /usr/local/bin/
COPY config.yaml /etc/url-redirector/
ENTRYPOINT ["url-redirector"]
```
```bash
docker build -t url-redirector .
docker run -d -p 80:80 url-redirector
```

## License
MIT
