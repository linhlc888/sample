## Introduction

This is sample service using [gokums-cli](https://github.com/gokums/cli) to generate boiler template.

There are two main funcions allows user login and register information.

## Clone and install dependency packages

```
mkdir -p $GOPATH/src/github.com/linhlc888/

cd $GOPATH/src/github.com/linhlc888/

git clone https://github.com/linhlc888/sample.git
cd sample
dep ensure
```

## Start service

```
cd src/service/user
docker-compose up
```

## Test functions

### Register new user

```
curl -X POST \
  http://localhost:9019/v1/user/register \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"fullname":"test user",
	"email": "test@gmail.com",
	"password":"123455"
	
}'
```
### Login with email and password

```
curl -X POST \
  http://localhost:9019/v1/user/login \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"email": "test@gmail.com",
	"password":"123455"
	
}'
```
