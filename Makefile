.PHONY: jwt-certs
jwt-certs:
	openssl genrsa | openssl pkcs8 -topk8 -v2 aes-128-ecb -out ./auth/certs/jwt.p8 -passout pass:
	openssl pkcs8 -in ./auth/certs/jwt.p8 -out ./auth/certs/jwt.pem -passin pass:
	openssl rsa -in ./auth/certs/jwt.pem -out ./auth/certs/jwt.key
	openssl rsa -in ./auth/certs/jwt.key -pubout -out ./auth/certs/jwt.key.pub

.PHONY: cp-jwt-certs
cp-jwt-certs:
	cp ./auth/certs/jwt.key.pub ./user/certs

.PHONY: install
install:
	export GO111MODULE=off; go get -v google.golang.org/grpc
	export GO111MODULE=off; go get -v github.com/golang/protobuf/protoc-gen-go
	export GO111MODULE=on; go get -v github.com/go-kit/kit
	export GO111MODULE=off; go get -v github.com/kujtimiihoxha/kit
	export GO111MODULE=off; go get -v github.com/99designs/gqlgen/cmd
	export GO111MODULE=off; go get -v github.com/vektah/dataloaden
	export GO111MODULE=off; go get -v -tags 'postgres' github.com/golang-migrate/migrate/cmd/migrate
	export GO111MODULE=off; go get -v github.com/thrasher-/gocryptotrader

.PHONY: proto
proto:
	cd ${GOPATH}/src/github.com/emurmotol/project/auth/pkg/grpc/pb; ./compile.sh
	cd ${GOPATH}/src/github.com/emurmotol/project/user/pkg/grpc/pb; ./compile.sh

.PHONY: build
build: jwt-certs cp-jwt-certs proto 
	go build -o ./auth/auth ./auth/cmd
	go build -o ./user/user ./user/cmd
	go build -o ./api/api ./api/server

.PHONY: test
test:
	go test -v ./auth/... -cover -count=1
	go test -v ./user/... -cover -count=1

.PHONY: docker-images
docker-images:
	docker build ./auth -t emurmotol/auth:latest
	docker build ./user -t emurmotol/user:latest
	docker build ./api -t emurmotol/api:latest

.PHONY: auth
auth:
	docker-compose -f ./auth/docker-compose.yml up

.PHONY: user
user:
	docker-compose -f ./user/docker-compose.yml up

.PHONY: api
api:
	docker-compose -f ./api/docker-compose.yml up

.PHONY: server
server:
	docker-compose -f ./server/docker-compose.yml up

.PHONY: user-seeder
user-seeder:
	go run ./user/cmd/seeder/main.go -database postgres://root:postgres@localhost:5433/project?sslmode=disable

.PHONY: seeder
seeder: user-seeder

.PHONY: docker-down
docker-down:
	docker-compose -f ./user/docker-compose.yml down
	docker-compose -f ./auth/docker-compose.yml down
	docker-compose -f ./api/docker-compose.yml down
	docker-compose -f ./server/docker-compose.yml down

.PHONY: docker-up
docker-up:
	docker-compose -f ./user/docker-compose.yml up -d
	docker-compose -f ./auth/docker-compose.yml up -d
	docker-compose -f ./api/docker-compose.yml up -d
	docker-compose -f ./server/docker-compose.yml up -d

SERVER_GOCRYPTOTRADER_DAEMON_CONTAINER_ID:=$(shell docker ps -aqf "ancestor=server_gocryptotrader_daemon")
.PHONY: gocryptotrader
gocryptotrader:
	docker cp ./server/gocryptotrader/config.json ${SERVER_GOCRYPTOTRADER_DAEMON_CONTAINER_ID}:/root/.gocryptotrader
	docker commit ${SERVER_GOCRYPTOTRADER_DAEMON_CONTAINER_ID} server_gocryptotrader_daemon
	docker rm -f ${SERVER_GOCRYPTOTRADER_DAEMON_CONTAINER_ID} gocryptotrader_web
	docker-compose -f ./server/docker-compose.yml up -d gocryptotrader_daemon gocryptotrader_web

.PHONY: migrate-up
migrate-up:
	migrate -verbose -source file://server/postgres/migrations -database postgres://root:postgres@localhost:5433/project?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	# migrate -verbose -source file://server/postgres/migrations -database postgres://root:postgres@localhost:5433/project?sslmode=disable force 20190304220339 down
	migrate -verbose -source file://server/postgres/migrations -database postgres://root:postgres@localhost:5433/project?sslmode=disable down

