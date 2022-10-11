# hub-mirror

使用 docker.io 或其他镜像服务来提供（但不限于） gcr.io、k8s.gcr.io、quay.io、ghcr.io 等国外镜像加速下载服务

> 为减少重复请求，合理利用资源，建议提前在 issues 搜索镜像是否已转换过
> 
> 示例：https://github.com/togettoyou/hub-mirror/issues?q=gcr.io%2Fgoogle-samples%2Fmicroservices-demo%2Femailservice%3Av0.3.5

# 原理

https://mp.weixin.qq.com/s/Vt0FRTx1PsoYFdLa0QZzWw

# 开始使用

## 方案一：白嫖我的，点个 Star ，直接提交 issues

要求：严格按照模板规范提交，参考： [成功案例](https://github.com/togettoyou/hub-mirror/issues/1)
，[失败案例](https://github.com/togettoyou/hub-mirror/issues/2)

限制：每次提交最多 11 个镜像地址

本人 Docker 账号有每日镜像拉取限额，请勿滥用

## 方案二：自己动手，丰衣足食，Fork 本项目，绑定你自己的 DockerHub 账号或其他镜像服务账号

1. 绑定账号

    - 如果要使用 DockerHub 的镜像服务

      在 `Settings`-`Secrets`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（你的 Docker 用户名）
      和 `DOCKER_TOKEN`（你的 Docker 密码） 两个 Secrets

    - 如果需要使用其他镜像服务,例如腾讯云、阿里云等

      在 `Settings`-`Secrets`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（你的其他镜像服务用户名）
      和 `DOCKER_TOKEN`（你的其他镜像服务密码）以及 `DOCKER_REPOSITORY` 三个 Secrets

      `DOCKER_REPOSITORY`配置例子

        - 腾讯云: `ccr.ccs.tencentyun.com/xxxxxx`
        - 阿里云: `registry.cn-hangzhou.aliyuncs.com/xxxxxx`
        - 等其他云...

2. 在 Fork 的项目中开启 `Settings`-`Options`-`Features` 中的 `Issues` 功能

3. 在 `Issues`-`Labels` 选择 `New label` 依次添加三个 label ：`hub-mirror`、`success`、`failure`

4. 在 `Actions` 里选择 `hub-mirror` ，在右边 `···` 菜单里选择 `Enable Workflow`

## 方案三：已有魔法，支持本地使用

```shell
$ go install github.com/togettoyou/hub-mirror@latest
```

```shell
$ hub-mirror --username=xxxxxx --password=xxxxxx --content='{ "hub-mirror": ["gcr.io/google-samples/microservices-demo/emailservice:v0.3.5","hello-world:latest"] }'
# 如果需要使用自定义镜像仓库
$ hub-mirror --username=xxxxxx --password=xxxxxx --repository=registry.cn-hangzhou.aliyuncs.com/xxxxxx --content='{ "hub-mirror": ["gcr.io/google-samples/microservices-demo/emailservice:v0.3.5","hello-world:latest"] }'
```

