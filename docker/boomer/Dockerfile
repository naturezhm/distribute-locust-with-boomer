# docker基础镜像 - locust运行镜像
#
# date 2019-11-27
# author naturezhm
# version 1.0

# FROM docker.io/bradyzm/py-base:1.0
FROM  docker.io/bradyzm/py-base:1.0
LABEL maintainer="naturezhm@gmail.com"


ARG APP_VERSION=1.0

ENV APP_VERSION=${APP_VERSION}

# WORKDIR /app

# RUN pip3 install locust

# EXPOSE 5557 8089

ENTRYPOINT [ "tail", "-f" ]