# auth
Authentication Microservice

## Pre-Installation

Download and install: https://github.com/protocolbuffers/protobuf/releases

```bash
go get -u -v google.golang.org/grpc

go get -u -v github.com/golang/protobuf/protoc-gen-go

go get -u -v github.com/go-kit/kit

go get -u -v github.com/kujtimiihoxha/kit
```

Create a new service

```bash
kit new service auth_api

kit generate service auth_api -w --gorilla

kit generate service auth_api -w -t grpc
```

Setup go mod

```bash
cd auth_api

go mod init github.com/emurmotol/project/auth_api/cmd

go mod tidy
```


