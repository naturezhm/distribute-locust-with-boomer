# docker基础镜像 - locust运行镜像
#
# date 2019-11-27
# author naturezhm
# version 1.0

FROM golang:1.14.1

LABEL maintainer="naturezhm@gmail.com"


RUN apt-get install git
ENV GOPATH /app
ENV GO111MODULE on

WORKDIR /app/src/github.com/naturezhm/distribute-locust-with-boomer/docker/boomer
COPY . /app/src/github.com/naturezhm/distribute-locust-with-boomer/docker/boomer

RUN chmod 755 ./start.sh
RUN go build ./...
RUN go build -o slave . 
RUN chmod 755 ./slave

ENTRYPOINT ["./start.sh"]