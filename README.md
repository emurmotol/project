# auth
Authentication Microservice

## Bootstrap Steps

Download and install: https://github.com/protocolbuffers/protobuf/releases

```bash
go get -u -v google.golang.org/grpc

go get -u -v github.com/golang/protobuf/protoc-gen-go

go get -u -v github.com/go-kit/kit

go get -u -v github.com/kujtimiihoxha/kit
```

Create a new service

```bash
cd project

kit new service auth_api

kit generate service auth_api -w --gorilla

kit generate service auth_api -w -t grpc
```

Setup go mod

```bash
cd project

go mod init github.com/emurmotol/project

go mod init github.com/emurmotol/project/auth_api

cd auth_api

go mod tidy
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


