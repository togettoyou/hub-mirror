package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/pflag"
)

var (
	content    = pflag.StringP("content", "", "", "原始镜像，格式为：{ \"hub-mirror\": [] }")
	maxContent = pflag.IntP("maxContent", "", 10, "原始镜像个数限制")
	username   = pflag.StringP("username", "", "", "docker hub 用户名")
	password   = pflag.StringP("password", "", "", "docker hub 密码")
	outputPath = pflag.StringP("outputPath", "", "output.sh", "结果输出路径")
)

func main() {
	pflag.Parse()

	fmt.Println("验证原始镜像内容")
	var hubMirrors struct {
		Content []string `json:"hub-mirror"`
	}
	err := json.Unmarshal([]byte(*content), &hubMirrors)
	if err != nil {
		panic(err)
	}
	if len(hubMirrors.Content) > *maxContent {
		panic("content is too long.")
	}
	fmt.Printf("%+v\n", hubMirrors)

	fmt.Println("连接 Docker")
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Println("验证 Docker 用户名密码")
	if *username == "" || *password == "" {
		panic("username or password cannot be empty.")
	}
	authConfig := types.AuthConfig{
		Username: *username,
		Password: *password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	_, err = cli.RegistryLogin(context.Background(), authConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("开始转换镜像")
	output := make([]struct {
		Source string
		Target string
	}, 0)

	wg := sync.WaitGroup{}

	for _, source := range hubMirrors.Content {
		if source == "" {
			continue
		}

		target := *username + "/" + strings.ReplaceAll(source, "/", ".")

		wg.Add(1)
		go func(source, target string) {
			defer wg.Done()

			fmt.Println("开始转换", source, "=>", target)
			ctx := context.Background()

			// 拉取镜像
			pullOut, err := cli.ImagePull(ctx, source, types.ImagePullOptions{})
			if err != nil {
				panic(err)
			}
			defer pullOut.Close()
			io.Copy(os.Stdout, pullOut)

			// 重新标签
			err = cli.ImageTag(ctx, source, target)
			if err != nil {
				panic(err)
			}

			// 上传镜像
			pushOut, err := cli.ImagePush(ctx, target, types.ImagePushOptions{
				RegistryAuth: authStr,
			})
			if err != nil {
				panic(err)
			}
			defer pushOut.Close()
			io.Copy(os.Stdout, pushOut)

			output = append(output, struct {
				Source string
				Target string
			}{Source: source, Target: target})
			fmt.Println("转换成功", source, "=>", target)
		}(source, target)
	}

	wg.Wait()

	if len(output) == 0 {
		panic("output is empty.")
	}

	tmpl, err := template.New("pull_images").Parse(`{{- range . -}}

docker pull {{ .Target }}
docker tag {{ .Target }} {{ .Source }}

{{ end -}}`)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(*outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = tmpl.Execute(outputFile, output)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
