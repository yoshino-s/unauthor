# Base
FROM golang:1.20.5-alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o ./cmd/unauthor ./cmd/unauthor

# Release
FROM alpine:3.18.2
COPY --from=builder /app/cmd/unauthor/unauthor /usr/local/bin/

ENTRYPOINT ["unauthor"]
