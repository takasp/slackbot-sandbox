FROM golang:1.11.4-alpine3.8

WORKDIR /go/src/projects

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

ENV GO111MODULE=on

CMD ["fresh"]