---
name: hub-mirror issue template
about: 用于执行 hub-mirror workflow 的 issue 模板
title: "[hub-mirror] 请求执行任务"
labels: ["hub-mirror"]
---

{
    "hub-mirror": [
        "格式：你需要转换的镜像$自定义镜像名",
        "实例1：ghcr.io/jenkins-x/jx-boot:3.10.3",
        "实例2：ghcr.io/jenkins-x/jx-boot:3.10.3$jx-boot",
        "要求：hub-mirror 标签是必选的，标题随意，每次最多支持转换 11 个镜像",
        "建议：修改这个 json",
        "......"
    ]
}