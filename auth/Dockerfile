FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/project/auth
ADD . /go/src/github.com/emurmotol/project/auth

RUN go get -v github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/emurmotol/project/auth
ENV GO111MODULE=on
RUN go mod tidy

ENTRYPOINT watcher -run github.com/emurmotol/project/auth/cmd -watch github.com/emurmotol/project/auth
