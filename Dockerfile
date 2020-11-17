FROM golang:1.14-alpine3.11 as builder
LABEL "name"="movie.api"
LABEL "maintainer"="kimtaek <jinze1991@icloud.com>"
LABEL "version"="1.0.0"

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go env && adduser -D -g '' service && apk add --no-cache git
ADD . /go/src
WORKDIR /go/src
RUN go build -o build/main main.go \
 && go build -o build/healthz cmd/healthz/healthz.go
USER service

FROM alpine:3.11
WORKDIR /go/src
COPY --from=builder /go/src/build .
ENTRYPOINT ["/go/src/main"]
HEALTHCHECK --start-period=2s --interval=10s --timeout=5s CMD ["/go/src/healthz"]