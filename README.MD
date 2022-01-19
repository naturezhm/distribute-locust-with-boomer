# Distribute Load Test Service Based On Locust And Boomer
> 分布式 locust+boomer 压测服务


## Require


### Run on K8S
```bash
apply                          kubectl apply 指定应用到指定集群
delete                         kubectl delete 同上，卸载
buildApply                     编译-打包镜像-部署一键 (latest) 或者指定tag

eg: 
make apply app=boomer tag:v1
```

## Install
```bash
build                          编译docker镜像 会自动编译latest 和 指定tag
push                           push指定latest
pushTag                        push指定tag
```

## Get Start

```bash

make help  查看更多指引

```


