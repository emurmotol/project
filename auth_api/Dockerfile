FROM golang

RUN mkdir -p /go/src/github.com/emurmotol/project/auth_api
ADD . /go/src/github.com/emurmotol/project/auth_api

RUN go get -v github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN git config --global url."https://9cfb59a44a7de832b75490b8563e19448c17d44f:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/emurmotol/project/auth_api
ENV GO111MODULE=on
RUN go mod download

ENTRYPOINT watcher -run github.com/emurmotol/project/auth_api/cmd -watch github.com/emurmotol/project/auth_api
