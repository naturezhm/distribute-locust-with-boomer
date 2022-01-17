# docker基础镜像 - locust运行镜像
#
# date 2019-11-27
# author naturezhm
# version 1.0

# FROM docker.io/bradyzm/py-base:1.0
# FROM  docker.io/bradyzm/py-base:1.0
# LABEL maintainer="naturezhm@gmail.com"


# ARG APP_VERSION=1.0

# ENV APP_VERSION=${APP_VERSION}

# # WORKDIR /app

# # RUN pip3 install locust

# # EXPOSE 5557 8089

# ENTRYPOINT [ "tail", "-f" ]

FROM python:3.9

#install locust
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential libncursesw5-dev libreadline-dev libssl-dev libgdbm-dev libc6-dev libsqlite3-dev libxml2-dev libxslt-dev python python-dev python-setuptools && apt-get clean
RUN pip install locust pyzmq

COPY start.sh /start.sh
COPY ./locustfile.py /locustfile.py
RUN chmod 755 /start.sh

EXPOSE 5557 5558 8089
ENTRYPOINT ["/start.sh"]