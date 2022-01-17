#!/bin/bash
# date 2022-01-17
# author naturezhm
# version 1.0
# 集成编译指令，方便使用，比直接调Shell更好用

HOST=$(shell hostname)
# 统一镜像前置名称 eg: load-test
IMAGE_NAME=load-test
# docker repository url
# registry-1.docker.io
IMAGE_URL=docker.io
IMAGE_PATH=$(IMAGE_URL)/bradyzm
K8S_APPLY_PATH=bradyzm


.PHONY: check-app

ifndef tag
override tag = latest
endif

check-app:
ifeq (,$(filter $(app),boomer locust))
	$(error USAGE: make [command] app=[boomer|locust])
else
	@echo "app:${app}"	
endif 

help: ## 帮助.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: check-app ## 编译docker镜像 会自动编译latest 和 指定tag
	@docker build -t $(IMAGE_PATH)/$(IMAGE_NAME)-$(app):latest -t $(IMAGE_PATH)/$(IMAGE_NAME)-$(app):$(tag) $(app)/

push: check-app ## push指定latest
	@docker push $(IMAGE_PATH)/$(IMAGE_NAME)-$(app):latest

pushTag: check-app ## push指定tag
	@docker push $(IMAGE_PATH)/$(IMAGE_NAME)-$(app):$(tag)

# composeUp: check-app  ## docker-compose 启动
# 	$(shell HOST=$(HOST) docker-compose -f $(app)/docker-compose.yml up -d)

# composeDown: check-app ## docker-compose 停止
# 	$(shell HOST=$(HOST) docker-compose -f $(app)/docker-compose.yml down)

apply: check-app  ## kubectl apply 指定应用到指定集群
	@echo "$(shell DOCKER_IMAGE=$(K8S_APPLY_PATH)/$(IMAGE_NAME)-$(app):$(tag) ./deploy-k8s.sh $(app) apply)"

delete: check-app ## kubectl delete 同上，卸载
	@echo "$(shell DOCKER_IMAGE=$(K8S_APPLY_PATH)/$(IMAGE_NAME)-$(app):$(tag) ./deploy-k8s.sh $(app) delete)"

buildApply: check-app build push apply ## 编译-打包镜像-部署一键 (latest)

