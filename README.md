# Project

This is the Mono Repo Project

## Getting Started

- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [Services](#services)
- [Usage](#usage)
- [Postman Collection](#postman-collection)
- [Donations](#donations)

## Dependencies

Requirements
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [protobuf](https://github.com/protocolbuffers/protobuf/releases)

Installable via Makefile
- [gRPC-Go](https://github.com/grpc/grpc-go)
- [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go)
- [GoKit](https://gokit.io/)
- [GoKit CLI](https://github.com/kujtimiihoxha/kit)
- [gqlgen](https://github.com/99designs/gqlgen)
- [dataloaden](github.com/vektah/dataloaden)
- [postgres](https://www.postgresql.org/)

## Configuration

- GO111MODULE=on

## Services

- [API](https://github.com/emurmotol/project/tree/master/api) - API Microservice (GraphQL)
- [Auth API](https://github.com/emurmotol/project/tree/master/auth_api) - Authentication Microservice
- [User API](https://github.com/emurmotol/project/tree/master/user_api) - User Management Microservice

## Usage

A Makefile is included for convenience

Install the dependencies
```bash
make install
```

Build the binaries
```bash
make build
```

Run the service
```bash
./service_name/service_name
```

Build the docker images
```bash
make docker
```

Start a service
```bash
make service_name
```

Run tests
```bash
make test
```

Generated with
```bash
# create a new service
kit new service service_name

# add endpoints to the service

# http transport
kit generate service service_name -w --gorilla

# grpc transport
kit generate service service_name -w -t grpc
```

Setup go modules
```bash
cd service_name

go mod init

go mod tidy
```

## Postman Collection

Import link: https://www.getpostman.com/collections/c162d7e4484b3c51e985

## Donations

If you like it you can send a small donation to any of the following addresses:

BTC: `3P1eTCYEcFGoN4bCfRAUbedfqK17DCMM5R`

ETH: `0x4939e019c56a8885bcd5fac11eba1cb1b147dc6e`

XRP: `rU2mEJSLqBRkYLVTv55rFTgQajkLTnT6mA` DT: `110892`

BAT: `0x0a317eA88131eFD0FC48E0ac9945996Eb690dbc0`