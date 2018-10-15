FROM golang:latest
MAINTAINER <284077318@qq.com>
COPY . $GOPATH/src/github.com/jicg/liteblog
WORKDIR $GOPATH/src/github.com/jicg/liteblog
RUN go get  github.com/jicg/liteblog
RUN go install -a github.com/jicg/liteblog

FROM debian:latest
MAINTAINER <284077318@qq.com>
COPY --from=0 /go/bin/liteblog /usr/bin/liteblog
COPY --from=0 /go/src/github.com/jicg/liteblog/start.sh /app/start.sh
COPY --from=0 /go/src/github.com/jicg/liteblog/views /app/views
COPY --from=0 /go/src/github.com/jicg/liteblog/static /app/static
COPY --from=0 /go/src/github.com/jicg/liteblog/conf /app/conf

VOLUME /app/data
VOLUME /app/assert

EXPOSE 8080
WORKDIR /app

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo 'Asia/Shanghai' >/etc/timezone \
&& chmod +x start.sh
CMD ["./start.sh"]