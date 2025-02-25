## hub-mirror

使用国内镜像仓库来提供（但不限于） docker.io、gcr.io、registry.k8s.io、k8s.gcr.io、quay.io、ghcr.io 等国外镜像加速下载服务

示例：https://github.com/togettoyou/hub-mirror/issues/2816

<img src="https://github.com/user-attachments/assets/ea93572c-6c05-4751-bde7-35a58fe083f1" width="520" alt="gopher云原生公众号二维码">

👆 扫码或搜索关注公众号：**gopher云原生**

## 使用流程

### 1️⃣ Fork 本项目

`Fork` 该项目，后续所有操作都在你 `Fork` 的仓库中进行。

### 2️⃣ 绑定账号

- 进入 `Settings` → `Secrets and variables` → `Actions`
- 选择 `New repository secret`，并添加以下 `Secrets`：

  - `DOCKER_USERNAME`：镜像仓库登录名
  - `DOCKER_TOKEN`：镜像仓库密码
  - `DOCKER_REPOSITORY`：镜像仓库地址

其中 `DOCKER_REPOSITORY` 示例：

- 腾讯云: `ccr.ccs.tencentyun.com/[namespace]`
- 阿里云: `registry.cn-hangzhou.aliyuncs.com/[namespace]`

例如我的是：`registry.cn-hangzhou.aliyuncs.com/hubmirrorbytogettoyou`

🔹 **示例截图**  

![阿里云镜像仓库](https://github.com/user-attachments/assets/6d7f3fda-cc8c-40dd-adf8-627a704c8533)

![Secrets 配置示例](https://github.com/user-attachments/assets/13010521-13b2-4c55-83d6-50956e039434)


### 3️⃣ 开启 Issues 功能

- 进入 `Settings` → `General` → `Features`
- 启用 `Issues`

🔹 **示例截图**  

![开启 Issues](https://github.com/user-attachments/assets/f981a0b9-b164-4582-8f5e-46d8cbe41bae)


### 4️⃣ 配置 Actions 权限

- 进入 `Settings` → `Actions` → `General`
- 在 `Workflow permissions` 选项中，选择：
  - ✅ `Read and write permissions`

🔹 **示例截图**  

![修改 Actions 权限](https://github.com/user-attachments/assets/9f556ced-d134-41f7-b47e-fa95c10db08a)


### 5️⃣ 添加 Issue Labels

- 进入 `Issues` → `Labels`
- 点击 `New label`
- 依次添加以下 Labels：

  - `hub-mirror`
  - `success`
  - `failure`

🔹 **示例截图**  

![添加 Labels](https://github.com/user-attachments/assets/b03db5eb-2401-49ce-ad12-515969dec27d)


### 6️⃣ 启用 Actions Workflow

- 进入 `Actions`
- 选择 `hub-mirror`
- 在右上角 `···` 菜单中选择 `Enable Workflow`

🔹 **示例截图**  

![启用 Workflow](https://github.com/user-attachments/assets/0709ac59-a731-4266-826e-0c619e933853)


### 7️⃣ 提交 Issue 触发同步

- 在 `Fork` 的仓库 `Issues` 页面，点击 `New issue`
- 选择 `hub-mirror` 模板，填写所需信息并提交

🔹 **示例截图**  

![提交 Issue](https://github.com/user-attachments/assets/c0357521-6dd0-4f13-8a99-bccdf1314ab8)

