# hub-mirror

使用 docker.io 来提供（但不限于） gcr.io、k8s.gcr.io、quay.io、ghcr.io 等国外镜像加速下载服务

# 使用

## 1. 白嫖我的，点个 Star ，直接提交 issues

要求：严格按照模板规范提交，参考： [成功案例](https://github.com/togettoyou/hub-mirror/issues/1) ，[失败案例](https://github.com/togettoyou/hub-mirror/issues/2)

限制：每次提交最多 11 个镜像地址

我的个人 Docker 账号有每日镜像拉取限额，请勿滥用

## 2. 自己动手，Fork 本项目，开启 issues 并绑定你自己的 Docker 账号

开启 `Settings`-`Options`-`Features` 中的 `Issues` 功能

在 `Settings`-`Secrets` 新建 `DOCKERHUB_USERNAME`（你的 Docker 用户名） 和 `DOCKERHUB_TOKEN`（你的 Docker 密码） 两个 Secrets

在 `Issues`-`Labels` 添加三个 label ：`hub-mirror`、`success`、`failure`

最后在 `Actions` 里选择 `hub-mirror` ，在右边 `···` 菜单里选择 `Enable Workflow`

## 3. 已有魔法，本地使用

```shell
go install github.com/togettoyou/hub-mirror@latest
```

```shell
hub-mirror --username=xxxxxx --password=xxxxxx --content='{ "hub-mirror": ["gcr.io/google-samples/microservices-demo/emailservice:v0.3.5"] }'
```

# 教程

https://mp.weixin.qq.com/s/Vt0FRTx1PsoYFdLa0QZzWw
