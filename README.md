## hub-mirror

使用 docker.io (hub.docker.com) 或其它国内镜像仓库来提供（但不限于） gcr.io、registry.k8s.io、k8s.gcr.io、quay.io、ghcr.io
等国外镜像加速下载服务

最近更新：

- 转换脚本提供 `docker`、`ctr`（containerd）命令

为减少重复请求，合理利用资源，建议提前在 issues 中搜索镜像是否已转换过

示例：[issues搜索registry.k8s.io/kube-apiserver:v1.27.4](https://github.com/togettoyou/hub-mirror/issues?q=registry.k8s.io%2Fkube-apiserver%3Av1.27.4)

## 原理

[无法拉取 gcr.io 镜像？用魔法来打败魔法](https://mp.weixin.qq.com/s/Vt0FRTx1PsoYFdLa0QZzWw)

更多云原生技术可关注微信公众号：【gopher云原生】

<img src="https://user-images.githubusercontent.com/55381228/221747734-13783ce6-1969-4c10-acd6-833f5046aa85.png" width="300px">

## 开始使用

#### 用法一：白嫖我的，点个 Star ，[直接提交 issues ](https://github.com/togettoyou/hub-mirror/issues/new/choose)，按照 issue 模板修改内容即可

要求：严格按照模板规范提交，参考： [成功案例](https://github.com/togettoyou/hub-mirror/issues/1813)

> 当任务失败时，可以查看失败原因并直接修改 issues 的内容，即可重新触发任务执行

> 限制：每次提交最多 11 个镜像地址（为啥是11个？因为我的第一次需求刚好要转换11个镜像🤣）

> 本人 Docker 账号有每日镜像拉取限额，请勿滥用

#### 用法二：自己动手，Fork 本项目，绑定你自己的 DockerHub 账号或其他镜像服务账号

1. 绑定账号

    - 如果要使用默认的 hub.docker.com 镜像服务

      在 `Settings`-`Secrets and variables`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（你的 Docker
      用户名） 和 `DOCKER_TOKEN`（你的 Docker 密码） 两个 Secrets

    - 如果需要使用其它镜像服务，例如腾讯云、阿里云等

      在 `Settings`-`Secrets and variables`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（你的其它镜像服务用户名）
      和 `DOCKER_TOKEN`（你的其它镜像服务密码）以及 `DOCKER_REPOSITORY` 三个 Secrets

      其中 `DOCKER_REPOSITORY` 配置例子：

        - 腾讯云: `ccr.ccs.tencentyun.com/xxxxxx`
        - 阿里云: `registry.cn-hangzhou.aliyuncs.com/xxxxxx`

2. 在 Fork 的项目中开启 `Settings`-`General`-`Features` 中的 `Issues` 功能

3. 在 Fork 的项目中修改 `Settings`-`Actions`-`General` 中的 `Workflow permissions` 为 `Read and write permissions`

4. 在 `Issues`-`Labels` 选择 `New label` 依次添加三个 label ：`hub-mirror`、`success`、`failure`

5. 在 `Actions` 里选择 `hub-mirror` ，在右边 `···` 菜单里选择 `Enable Workflow`

6. 在 Fork 的项目中提交 issues
