## hub-mirror

使用国内镜像仓库来提供（但不限于） docker.io、gcr.io、registry.k8s.io、k8s.gcr.io、quay.io、ghcr.io 等国外镜像加速下载服务

示例：https://github.com/togettoyou/hub-mirror/issues/2816

## 开始使用

[提交 issues](https://github.com/togettoyou/hub-mirror/issues/new/choose)，按照 issue 模板内容修改即可

<img src="https://github.com/user-attachments/assets/ea93572c-6c05-4751-bde7-35a58fe083f1" width="520" alt="gopher云原生公众号二维码">

👆 扫码或搜索关注公众号：**gopher云原生**

## 具体使用流程

#### Fork 本项目（可设为私有），绑定你自己的国内镜像仓库

1. 绑定账号

   在 `Settings`-`Secrets and variables`-`Actions` 选择 `New repository secret` 新建 `DOCKER_USERNAME`（镜像仓库登录名）
   和 `DOCKER_TOKEN`（镜像仓库密码）以及 `DOCKER_REPOSITORY` 三个 Secrets

   ![image](https://github.com/user-attachments/assets/13010521-13b2-4c55-83d6-50956e039434)

   其中 `DOCKER_REPOSITORY` 配置例子：

    - 腾讯云: `ccr.ccs.tencentyun.com/[namespace]`
    - 阿里云: `registry.cn-hangzhou.aliyuncs.com/[namespace]`

   例如我的是：`registry.cn-hangzhou.aliyuncs.com/hubmirrorbytogettoyou`

   ![image](https://github.com/user-attachments/assets/5af044b7-f62e-401c-976f-a8556964b995)

2. 开启 `Settings`-`General`-`Features` 中的 `Issues` 功能

   ![image](https://github.com/user-attachments/assets/f981a0b9-b164-4582-8f5e-46d8cbe41bae)

3. 修改 `Settings`-`Actions`-`General` 中的 `Workflow permissions` 为 `Read and write permissions`

   ![image](https://github.com/user-attachments/assets/9f556ced-d134-41f7-b47e-fa95c10db08a)

4. 在 `Issues`-`Labels` 选择 `New label` 依次添加三个 label ：`hub-mirror`、`success`、`failure`

   ![image](https://github.com/user-attachments/assets/b03db5eb-2401-49ce-ad12-515969dec27d)

5. 在 `Actions` 里选择 `hub-mirror` ，在右边 `···` 菜单里选择 `Enable Workflow`

   ![image](https://github.com/user-attachments/assets/0709ac59-a731-4266-826e-0c619e933853)

6. 在 Fork 的项目中提交 issues

   ![image](https://github.com/user-attachments/assets/c0357521-6dd0-4f13-8a99-bccdf1314ab8)
