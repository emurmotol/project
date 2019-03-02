# Project

This is the Mono Repo Project

Generated with

```bash
kit new service service_name

# http transport

kit generate service service_name -w --gorilla

# grpc transport

kit generate service service_name -w -t grpc
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- TODO: todo

## Dependencies

Download and install: https://github.com/protocolbuffers/protobuf/releases

```
TODO
```

## Usage

A Makefile is included for convenience

Install the dependencies
```
make
```

Build the binaries
```
make build
```

Run the service
```
./service_name/service_name
```

Build a docker image
```
make docker
```

Run start the services
```
make docker-compose
```

Run tests
```
make test
```