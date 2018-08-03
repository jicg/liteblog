FROM golang:latest

MAINTAINER <284077318@qq.com>

COPY . $GOPATH/src/github.com/jicg/liteblog

WORKDIR $GOPATH/src/github.com/jicg/liteblog

WORKDIR $GOPATH/src/github.com/jicg/liteblog

RUN go get  github.com/jicg/liteblog

RUN go install -a github.com/jicg/liteblog
