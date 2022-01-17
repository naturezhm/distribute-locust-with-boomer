#!/bin/bash
# date 2021-11-18
# author naturezhm
# version 1.0
#
# THis shell script used for how to build image.

isLatest=$1

IMAGE_VERSION=1.0
APP_VERSION=1.0
IMAGE_NAME=locust-master
# docker repository url
# registry-1.docker.io
IMAGE_URL=docker.io
IMAGE_PATH=$IMAGE_URL/bradyzm

if [ "x$isLatest" = "x" ] 
then  
isLatest=0
fi

docker build -t $IMAGE_PATH/$IMAGE_NAME:$IMAGE_VERSION --build-arg APP_VERSION=$APP_VERSION .


if [ "$?" != 0 ] ; then
    printf "build faild !!! \n" ;
    exit 1;
fi


printf "======  Please Login to ${IMAGE_URL}  ====== \n"
docker login $IMAGE_URL


if [ "$?" != 0 ] ; then
    printf "Login faild to ${IMAGE_URL}!!! \n" ;
    exit 1;
fi

docker push $IMAGE_PATH/$IMAGE_NAME:$IMAGE_VERSION

if [ "$?" != 0 ] ; then
    printf "Push faild!!! \n" ;
    exit 1;
fi

if [ "$isLatest" == "latest" ]; then 
    docker tag $IMAGE_PATH/$IMAGE_NAME:$IMAGE_VERSION $IMAGE_PATH/$IMAGE_NAME:latest
    docker push $IMAGE_PATH/$IMAGE_NAME:latest

    if [ "$?" != 0 ] ; then
         printf "Push faild!!! \n" ;
        exit 1;
    fi
printf "====== Finish push $IMAGE_PATH/$IMAGE_NAME:latest  ====== \n"
fi

printf "====== Finish push $IMAGE_PATH/$IMAGE_NAME:$IMAGE_VERSION ====== \n"
