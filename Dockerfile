FROM golang:1.14-alpine3.11 as builder
LABEL "name"="movie.api"
LABEL "maintainer"="kimtaek <jinze1991@icloud.com>"
LABEL "version"="1.0.0"

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl

WORKDIR /go/app

COPY . .
RUN go build -o build/main main.go \
 && go build -o build/healthz cmd/healthz/healthz.go

FROM alpine:3.11
WORKDIR /app
COPY --from=builder /go/app/build .
RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app

ENTRYPOINT ["/app/main"]
HEALTHCHECK --start-period=2s --interval=10s --timeout=5s CMD ["/app/healthz"]