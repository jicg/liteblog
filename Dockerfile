FROM golang:latest
MAINTAINER <284077318@qq.com>
COPY . $GOPATH/src/github.com/jicg/liteblog
WORKDIR $GOPATH/src/github.com/jicg/liteblog
RUN go get  github.com/jicg/liteblog
RUN go install -a github.com/jicg/liteblog

FROM debian:latest
MAINTAINER <284077318@qq.com>
COPY --from=0 /go/bin/easypos /usr/bin/liteblog
COPY --from=0 /go/src/github.com/jicg/liteblog/views /app/views
COPY --from=0 /go/src/github.com/jicg/liteblog/static /app/static
COPY --from=0 /go/src/github.com/jicg/liteblog/conf /app/conf
VOLUME /app/data
VOLUME /app/log
VOLUME /app/conf
EXPOSE 8080
WORKDIR /app
CMD /usr/bin/liteblog