# Project

This is the Mono Repo Project

## Getting Started

- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [Services](#services)
- [Usage](#usage)
- [Notes](#notes)
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
- [redis](https://redis.io/)
- [migrate](https://github.com/golang-migrate/migrate/tree/master/cli)
- [gocryptotrader](https://github.com/thrasher-/gocryptotrader)

## Configuration

- GO111MODULE=on

## Services

- [API](https://github.com/emurmotol/project/tree/master/api) - API Microservice (GraphQL)
- [Auth](https://github.com/emurmotol/project/tree/master/auth) - Authentication Microservice
- [User](https://github.com/emurmotol/project/tree/master/user) - User Management Microservice

## Usage

A Makefile is included for convenience

Install the dependencies
```bash
make install
```

Run migrations
```bash
make migrate-up
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
make docker-images
```

Start a service
```bash
make service_name

# or start all services
make docker-up
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

kit generate service service_name -w -t grpc
```

Setup go modules
```bash
cd service_name

go mod init

go mod tidy
```

Create migration
```bash
migrate create -ext sql -dir=./server/postgres/migrations create_sample_table
```

## Notes

```bash
$(ifconfig enp2s0 | sed -En -e 's/.*inet ([0-9.]+).*/\1/p')
```

## Donations

If you like it you can send a small donation to any of the following addresses:

BTC: `3P1eTCYEcFGoN4bCfRAUbedfqK17DCMM5R`

ETH: `0x4939e019c56a8885bcd5fac11eba1cb1b147dc6e`

XRP: `rU2mEJSLqBRkYLVTv55rFTgQajkLTnT6mA` DT: `110892`

BAT: `0x0a317eA88131eFD0FC48E0ac9945996Eb690dbc0`