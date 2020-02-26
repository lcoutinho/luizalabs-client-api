FROM golang:1.11-alpine as builder

RUN grep nobody /etc/passwd > /etc/passwd.nobody \
    && grep nobody /etc/group > /etc/group.nobody
RUN apk --no-cache update \
    && apk add --no-cache ca-certificates git \
    && wget -O- https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/github.com/lcoutinho/luizalabs-client-api

EXPOSE 8000