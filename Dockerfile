FROM golang:latest as builder
MAINTAINER <284077319@qq.com>
WORKDIR /go/src/github.com/jicg/liteblog
RUN go get github.com/tools/godep
COPY . .
#ENV CGO_ENABLED=0
#RUN godep go build -installsuffix cgo -ldflags="-w -s"
RUN godep go build  -ldflags="-w -s"
#FROM scratch as final
#FROM alpine as final
FROM debian:latest as final
MAINTAINER <284077319@qq.com>

#COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder go/src/github.com/jicg/liteblog/liteblog /app/liteblog
COPY --from=builder go/src/github.com/jicg/liteblog/views /app/views
COPY --from=builder go/src/github.com/jicg/liteblog/static /app/static
COPY --from=builder go/src/github.com/jicg/liteblog/conf /app/conf

VOLUME /app/data
VOLUME /app/assert
EXPOSE 8080
WORKDIR /app
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo 'Asia/Shanghai' >/etc/timezone \
#&& chmod +x start.sh
#RUN chmod +x liteblog
ENTRYPOINT ["/app/liteblog"]
CMD []