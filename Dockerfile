# Build golang image
FROM golang:1.19.4-alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN go build -o /http-server

FROM alpine:3.9

COPY --from=builder /http-server /app/http-server

ENTRYPOINT ["/app/http-server"]
