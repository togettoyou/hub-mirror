---
name: hub-mirror issue template
about: 用于执行 hub-mirror workflow 的 issue 模板
title: "[hub-mirror] 请求执行任务"
labels: ["hub-mirror"]
---

{
    "hub-mirror": [
        "格式：你需要转换的镜像$自定义镜像名:自定义标签名",
        "$自定义镜像名:自定义标签名 是可选的",
        "示例1：registry.k8s.io/kube-apiserver:v1.27.4",
        "示例2：registry.k8s.io/kube-apiserver:v1.27.4$my-kube-apiserver",
        "示例3：registry.k8s.io/kube-apiserver:v1.27.4$my-kube-apiserver:mytag",
        "要求：hub-mirror 标签是必选的，标题随意，每次最多支持转换 11 个镜像",
        "建议：按照这个 json 格式修改",
        "请确保 json 格式是正确的，比如这里最后一个是没有逗号的"
    ]
}