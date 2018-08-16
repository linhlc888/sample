FROM golang:1.9-alpine3.6

RUN apk --no-cache add git ca-certificates tzdata && update-ca-certificates

## Create gokums user
#RUN addgroup -S gokums &&   adduser -S gokums -G gokums &&   mkdir -p /gokums
#RUN addgroup -S gokums &&   adduser -S gokums -G gokums
#USER gokums
#WORKDIR /gokums
#ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-386 /usr/bin/dep
RUN go get -u github.com/golang/dep/cmd/dep

#COPY dist/user /gokums/user

#RUN chown -R gokums:gokums /gokums

#RUN mkdir -p $GOPATH/src/user_service
WORKDIR $GOPATH/src/github.com/linhlc888/sample/src/service/user
COPY . ./
RUN dep init && dep ensure --vendor-only && go build -o .build/app .
EXPOSE 9019
ENTRYPOINT [".build/app"]
