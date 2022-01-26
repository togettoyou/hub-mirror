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

教程首发微信公众号：【寻寻觅觅的Gopher】，欢迎关注

![微信公众号.png](https://cdn.nlark.com/yuque/0/2021/png/1077776/1628483947581-9a649b2f-a0bb-4ef4-879d-92ab6e9fddde.png)

目前常用的 Docker Registry 公开服务有：

- `docker.io` ：Docker Hub 官方镜像仓库，也是 Docker 默认的仓库

- `gcr.io`、`k8s.gcr.io` ：谷歌镜像仓库

- `quay.io` ：Red Hat 镜像仓库

- `ghcr.io` ：GitHub 镜像仓库

当使用 `docker pull 仓库地址/用户名/仓库名:标签` 时，会前往对应的仓库地址拉取镜像，标签无声明时默认为 `latest`， 仓库地址无声明时默认为 `docker.io` 。

![](https://gitee.com/togettoyou/picture/raw/master/2022-1-24/1642987913348-Snipaste_2022-01-24_09-26-33.gif)

众所周知的原因，在国内访问这些服务异常的慢，甚至 `gcr.io` 和 `quay.io` 根本无法访问。

![](https://gitee.com/togettoyou/picture/raw/master/2022-1-24/1642994022304-carbon%20(5).gif)


## 解决方案：镜像加速器

针对 `Docker Hub` ，Docker 官方和国内各大云服务商均提供了 Docker 镜像加速服务。

你只需要简单配置一下（以 Linux 为例）：

```bash
sudo mkdir -p /etc/docker

sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["镜像加速器"]
}
EOF

sudo systemctl daemon-reload
sudo service docker restart
```

便可以通过访问国内镜像加速器来加速 `Docker Hub` 的镜像下载。

![](https://gitee.com/togettoyou/picture/raw/master/2022-1-24/1642989154674-Snipaste_2022-01-24_09-27-33.gif)

不过这种办法也只能针对 `docker.io` ，其它的仓库地址并没有真正实际可用的加速器（至少我目前没找到）。

## 解决方案：用魔法打败魔法

既然无法治本，那治治标还是可以的吧。

若我们使用一台魔法机器从 `gcr.io` 或 `quay.io` 等仓库先把我们无法下载的镜像拉取下来，然后重新上传到 `docker.io` ，是不是就可以使用 `Docker Hub` 的镜像加速器下载了。

![](https://gitee.com/togettoyou/picture/raw/master/2022-1-24/1642990777364-Snipaste_2022-01-24_09-28-51.gif)

镜像仓库迁移的功能，我这里采用了 Go Docker SDK ，整体实现也比较简单，参考 [main.go](https://github.com/togettoyou/hub-mirror/blob/main/main.go)

以需要转换的 `gcr.io/google-samples/microservices-demo/emailservice:v0.3.5` 为例，使用方式：

![](https://gitee.com/togettoyou/picture/raw/master/2022-1-24/1642991865261-carbon-_3_.gif)

功能实现了，剩下的就是找台带有魔法的机器了。

GitHub Actions 就是个好选择，我们可以利用提交 `issues` 来触发镜像仓库迁移的功能。

`workflow` 的实现参考 [hub-mirror.yml](https://github.com/togettoyou/hub-mirror/blob/main/.github/workflows/hub-mirror.yml)

实际的使用效果参考 [issues](https://github.com/togettoyou/hub-mirror/issues?q=is%3Aissue+is%3Aopen+label%3Ahub-mirror) 

只要执行最终输出的命令，就可以飞快的使用 Docker Hub 的加速器下载 `gcr.io` 或 `quay.io` 等镜像了。
