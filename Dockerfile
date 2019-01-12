FROM golang:latest as builder
MAINTAINER <284077319@qq.com>
WORKDIR $GOPATH/src/github.com/jicg/liteblog
COPY . .
RUN go get
RUN go install -a -ldflags="-w -s"

FROM scratch as final
MAINTAINER <284077319@qq.com>
WORKDIR /app
COPY --from=builder go/bin/liteblog liteblog
COPY --from=builder go/src/github.com/jicg/liteblog/views views
COPY --from=builder go/src/github.com/jicg/liteblog/static static
COPY --from=builder go/src/github.com/jicg/liteblog/conf conf

VOLUME /app/data
VOLUME /app/assert
EXPOSE 8080

#RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#&& echo 'Asia/Shanghai' >/etc/timezone \
#&& chmod +x start.sh
CMD ["./liteblog"]