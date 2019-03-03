FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/project/user_api
ADD . /go/src/github.com/emurmotol/project/user_api

RUN go get -v github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/emurmotol/project/user_api
ENV GO111MODULE=on
RUN go mod download

ENTRYPOINT watcher -run github.com/emurmotol/project/user_api/cmd -watch github.com/emurmotol/project/user_api
