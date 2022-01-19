# Distribute Load Test Service Based On Locust And Boomer
> 分布式 locust+boomer 压测服务


## Require

- k8s 运行环境 （本机已经配置好对应的kubecofig文件）
- 一个能开发go的IDE 推荐 vscode goland
- git
- docker

### Run on K8S
```bash
make apply                          kubectl apply 指定应用到指定集群
make delete                         kubectl delete 同上，卸载
make buildApply                     编译-打包镜像-部署一键 (latest) 或者指定tag


```

## Install
```bash
make build                          编译docker镜像 会自动编译latest 和 指定tag
make push                           push指定latest
make pushTag                        push指定tag
```

可以通过 `make help` 查看更多指引

## Get Start

```bash


git clone ssh://git@gitlab.internal.supermonkey.com.cn:10022/zhangming/distribute-locust-with-boomer.git
cd distribute-locust-with-boomer/docker
// 本地编译镜像
make build app=boomer tag:v1
// push镜像
make pushTag app=boomer tag:v1
// 应用镜像到集群 (确保已经安装好k8s配置文件)
make apply app=boomer tag:v1

```


## Develop
### 如何写压测脚本?


可以查看 `docker/locust/k8s.yml` 默认是只开启了 fastHttp-task 是对一个图片的压测，如果正式使用，需要先学习一下go，
至少需要知道go语言的 如何读 csv文件作为账号数据源

然后在 `docker/boomer/task` 目录编辑任务脚本



