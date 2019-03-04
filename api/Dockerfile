FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/project/api
ADD . /go/src/github.com/emurmotol/project/api

RUN go get -v github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/emurmotol/project/api
ENV GO111MODULE=on
RUN go mod tidy

ENTRYPOINT watcher -run github.com/emurmotol/project/api/server -watch github.com/emurmotol/project/api
