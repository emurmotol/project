# project

Mono Repo Project

## Bootstrap Steps

Download and install: https://github.com/protocolbuffers/protobuf/releases

```bash
go get -u -v google.golang.org/grpc

go get -u -v github.com/golang/protobuf/protoc-gen-go

go get -u -v github.com/go-kit/kit

go get -u -v github.com/kujtimiihoxha/kit
```

Repository

```
# create repo in github then clone it:

mkdir -p $GOPATH/src/github.com/emurmotol

cd $GOPATH/src/github.com/emurmotol

git clone git@github.com:emurmotol/project.git
```

Create a new service

```bash
cd project

kit new service service_name

# http transport

kit generate service service_name -w --gorilla

# grpc transport

kit generate service service_name -w -t grpc
```

Setup go modules

```bash
cd project

go mod init

cd service_name

go mod init

go mod tidy
```

## Postman Collection

Import link: https://www.getpostman.com/collections/c162d7e4484b3c51e985

## Microservices

### auth_api - Authentication microservice

Docker

```
docker pull emurmotol/auth_api:latest

docker run --rm -it -p 8081:8081 -p 8082:8082 emurmotol/auth_api:latest

# or

cd project

docker-compose up auth_api
```

Generate JWT certificates

```bash
cd auth_api

mkdir certs

openssl genrsa | openssl pkcs8 -topk8 -v2 aes-128-ecb -out certs/jwt.p8 -passout pass:

openssl pkcs8 -in certs/jwt.p8 -out certs/jwt.pem -passin pass:

openssl rsa -in certs/jwt.pem -out certs/jwt.key

openssl rsa -in certs/jwt.key -pubout -out certs/jwt.key.pub
```

gRPC client

proto: https://github.com/emurmotol/project/tree/master/auth_api/pkg/grpc/pb

HTTP endpoints

Included in postman collection provided above

location: (project/auth_api)

