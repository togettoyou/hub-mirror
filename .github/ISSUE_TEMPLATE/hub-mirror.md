---
name: hub-mirror issue template
about: 用于执行 hub-mirror workflow 的 issue 模板
title: "[hub-mirror] 请求执行任务"
labels: ["hub-mirror"]
---

{
    "platform":"",
    "hub-mirror": [
        "若对OS/ARCH无要求，platform请留空，不要加任何值，默认就是linux/amd64",
        "如需切换arm架构，请修改platform为arm64或linux/arm64/v8",
        "格式：你需要转换的原始镜像$自定义镜像名:自定义标签名 （其中 $自定义镜像名:自定义标签名 是可选的）",
        "以下是三个正确示例",
        "registry.k8s.io/kube-apiserver:v1.27.4",
        "registry.k8s.io/kube-apiserver:v1.27.4$my-kube-apiserver",
        "registry.k8s.io/kube-apiserver:v1.27.4$my-kube-apiserver:mytag",
        "要求：hub-mirror 标签是必选的，标题随意，内容严格按照该 json 格式，每次最多支持转换 11 个镜像",
        "错误的镜像都会被跳过",
        "请确保 json 格式是正确的，比如下面这个是最后一个，后面是没有逗号的",
        "好了，改成你想要转换的镜像吧"
    ]
}
