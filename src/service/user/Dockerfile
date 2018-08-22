FROM alpine:3.6

RUN apk --no-cache add git ca-certificates tzdata && update-ca-certificates

## Create gokums user
RUN addgroup -S gokums &&   adduser -S gokums -G gokums &&   mkdir -p /gokums
RUN chown -R gokums:gokums /gokums
USER gokums
WORKDIR /gokums

COPY dist/app /gokums/app
EXPOSE 9019
ENTRYPOINT ["/gokums/app"]
