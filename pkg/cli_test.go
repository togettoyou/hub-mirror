package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource2Target(t *testing.T) {
	cli := &Cli{
		username:   "togettoyou",
		repository: "",
	}

	output, err := cli.Source2Target("")
	assert.Nil(t, output)

	source := "registry.k8s.io/kube-apiserver"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, source, output.Source)
	assert.Equal(t, "togettoyou/registry.k8s.io.kube-apiserver", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, source, output.Source)
	assert.Equal(t, "togettoyou/registry.k8s.io.kube-apiserver:v1.27.4", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4$kube-apiserver"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver:v1.27.4", output.Source)
	assert.Equal(t, "togettoyou/kube-apiserver:v1.27.4", output.Target)

	source = "registry.k8s.io/kube-apiserver:v1.27.4$kube-apiserver:mytag"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, "registry.k8s.io/kube-apiserver:v1.27.4", output.Source)
	assert.Equal(t, "togettoyou/kube-apiserver:mytag", output.Target)

	source = "nginx@sha256:123456$nginx"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, "nginx@sha256:123456", output.Source)
	assert.Equal(t, "togettoyou/nginx", output.Target)

	source = "nginx@sha256:123456$nginx:mytag"
	output, err = cli.Source2Target(source)
	assert.Nil(t, err)
	assert.Equal(t, "nginx@sha256:123456", output.Source)
	assert.Equal(t, "togettoyou/nginx:mytag", output.Target)
}
