#!/bin/sh

FILE_NAME='dist/app'

docker run \
  --rm \
  -it \
  -v $GOPATH/src/github.com/linhlc888/sample:/go/src/github.com/linhlc888/sample \
  --workdir /go/src/github.com/linhlc888/sample/src/service/user \
  golang:1.9-alpine3.6 \
  /bin/sh -c " \
    apk update; \
    go build -o ${FILE_NAME} .; \
  "