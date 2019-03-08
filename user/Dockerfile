FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/project/user
ADD . /go/src/github.com/emurmotol/project/user

RUN go get -v github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/emurmotol/project/user
ENV GO111MODULE=on
RUN go mod tidy

ENTRYPOINT watcher -run github.com/emurmotol/project/user/cmd -watch github.com/emurmotol/project/user
