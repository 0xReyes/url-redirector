FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o url-redirector main.go

FROM alpine
COPY --from=builder /app/url-redirector /usr/local/bin/
COPY config.yaml /etc/url-redirector/
ENTRYPOINT ["url-redirector"]
