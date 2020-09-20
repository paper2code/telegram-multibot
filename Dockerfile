FROM golang:1-alpine as builder
MAINTAINER paper2code <contact@paper2code.com>

WORKDIR /opt/service
COPY . .

FROM alpine:3.12

RUN apk add --no-cache ca-certificates

CMD ["tg-multibot"]
