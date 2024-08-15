<img src="https://user-images.githubusercontent.com/55381228/221747734-13783ce6-1969-4c10-acd6-833f5046aa85.png" width="300px">

## hub-mirror

使用国内镜像仓库来提供（但不限于） docker.io、gcr.io、registry.k8s.io、k8s.gcr.io、quay.io、ghcr.io 等国外镜像加速下载服务

示例：https://github.com/togettoyou/hub-mirror/issues/2816

## 试用

[直接在我的项目中提交 issues ](https://github.com/togettoyou/hub-mirror/issues/new/choose)，按照 issue 模板修改内容即可

> 个人配置的是阿里云个人实例镜像仓库，仓库限额为 300 ，所以有可能出现上传或拉取失败的情况（本人会不定时清理历史仓库

## 开始使用

#### Fork 本项目（可设为私有），绑定你自己的国内镜像仓库

1. 绑定账号

   在 `Settings`-`Secrets and variables`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（镜像仓库登录名）
   和 `DOCKER_TOKEN`（镜像仓库密码）以及 `DOCKER_REPOSITORY` 三个 Secrets

   其中 `DOCKER_REPOSITORY` 配置例子：

     - 腾讯云: `ccr.ccs.tencentyun.com/[namespace]`
     - 阿里云: `registry.cn-hangzhou.aliyuncs.com/[namespace]`
   
   例如我的是：`registry.cn-hangzhou.aliyuncs.com/hubmirrorbytogettoyou`

2. 开启 `Settings`-`General`-`Features` 中的 `Issues` 功能

3. 修改 `Settings`-`Actions`-`General` 中的 `Workflow permissions` 为 `Read and write permissions`

4. 在 `Issues`-`Labels` 选择 `New label` 依次添加三个 label ：`hub-mirror`、`success`、`failure`

5. 在 `Actions` 里选择 `hub-mirror` ，在右边 `···` 菜单里选择 `Enable Workflow`

6. 在 Fork 的项目中提交 issues
