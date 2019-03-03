.PHONY: jwt-certs
jwt-certs:
	openssl genrsa | openssl pkcs8 -topk8 -v2 aes-128-ecb -out ./auth_api/certs/jwt.p8 -passout pass:
	openssl pkcs8 -in ./auth_api/certs/jwt.p8 -out ./auth_api/certs/jwt.pem -passin pass:
	openssl rsa -in ./auth_api/certs/jwt.pem -out ./auth_api/certs/jwt.key
	openssl rsa -in ./auth_api/certs/jwt.key -pubout -out ./auth_api/certs/jwt.key.pub

.PHONY: cp-jwt-certs
cp-jwt-certs:
	cp ./auth_api/certs/jwt.key.pub ./user_api/certs

.PHONY: install
install: jwt-certs
	export GO111MODULE=off; go get -v google.golang.org/grpc
	export GO111MODULE=off; go get -v github.com/golang/protobuf/protoc-gen-go
	export GO111MODULE=on; go get -v github.com/go-kit/kit
	export GO111MODULE=off; go get -v github.com/kujtimiihoxha/kit
	export GO111MODULE=off; go get -v github.com/99designs/gqlgen/cmd
	export GO111MODULE=off; go get -v github.com/vektah/dataloaden

.PHONY: proto
proto:
	cd ${GOPATH}/src/github.com/emurmotol/project/auth_api/pkg/grpc/pb; ./compile.sh
	cd ${GOPATH}/src/github.com/emurmotol/project/user_api/pkg/grpc/pb; ./compile.sh

.PHONY: build
build: cp-jwt-certs proto 
	go build -o ./auth_api/auth_api ./auth_api/cmd
	go build -o ./user_api/user_api ./user_api/cmd
	go build -o ./api/api ./api/server

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build ./auth_api -t emurmotol/auth_api:latest
	docker build ./user_api -t emurmotol/user_api:latest
	docker build ./api -t emurmotol/api:latest

.PHONY: auth_api
auth_api:
	docker-compose -f ./auth_api/docker-compose.yml up auth_api

.PHONY: user_api
user_api:
	docker-compose -f ./user_api/docker-compose.yml up user_api

.PHONY: api
api:
	docker-compose -f ./api/docker-compose.yml up api

.PHONY: server
server:
	mkdir -p ./server/postgres/data
	mkdir -p ./server/redis/data
	docker-compose -f ./server/docker-compose.yml up
