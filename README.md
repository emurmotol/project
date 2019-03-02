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
```bash
TODO
```

## Usage

A Makefile is included for convenience

Install the dependencies
```bash
make
```

Build the binaries
```bash
make build
```

Run the service
```bash
./service_name/service_name
```

Build a docker image
```bash
make docker
```

Run start the services
```bash
make docker-compose
```

Run tests
```bash
make test
```