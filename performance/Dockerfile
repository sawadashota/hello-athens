FROM golang:1.13rc1-buster

RUN set -ex && \
    apt-get update && \
    apt-get install -y time && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV GO111MODULE="on"

WORKDIR /go/src/github.com/gohugoio/hugo
RUN git clone --progress --depth 1 https://github.com/gohugoio/hugo.git .
