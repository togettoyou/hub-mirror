package pkg

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Cli struct {
	cli        *client.Client
	repository string
	username   string
	auth       string
	log        io.Writer
}

func NewCli(ctx context.Context, repository, username, password string, log io.Writer) (*Cli, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password cannot be empty")
	}

	authConfig := types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: repository,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	_, err = cli.RegistryLogin(ctx, authConfig)
	if err != nil {
		return nil, err
	}

	return &Cli{
		cli:        cli,
		repository: repository,
		username:   username,
		auth:       base64.URLEncoding.EncodeToString(encodedJSON),
		log:        log,
	}, nil
}

type Output struct {
	Source string
	Target string
}

func (c *Cli) Source2Target(source string) (*Output, error) {
	if source == "" {
		return nil, errors.New("source is nil")
	}

	target := source

	if strings.Contains(source, "$") {
		parts := strings.Split(source, "$")
		if len(parts) > 1 {
			source = parts[0]
			target = parts[1]
		}
	}

	if !strings.Contains(target, ":") &&
		!strings.Contains(source, "@sha256") && strings.Contains(source, ":") {
		parts := strings.Split(source, ":")
		if len(parts) > 1 {
			target += ":" + parts[1]
		}
	}

	if c.repository == "" {
		target = c.username + "/" + strings.ReplaceAll(target, "/", ".")
	} else {
		target = c.repository + "/" + strings.ReplaceAll(target, "/", ".")
	}

	return &Output{
		Source: source,
		Target: target,
	}, nil
}

func (c *Cli) PullTagPushImage(ctx context.Context, source string) (*Output, error) {
	output, err := c.Source2Target(source)
	if err != nil {
		return nil, err
	}

	err = c.PullImage(ctx, output.Source)
	if err != nil {
		return nil, err
	}

	err = c.cli.ImageTag(ctx, output.Source, output.Target)
	if err != nil {
		return nil, err
	}

	err = c.PushImage(ctx, output.Target)
	if err != nil {
		return nil, err
	}

	return output, nil
}

type errorMessage struct {
	Error string `json:"error"`
}

func (c *Cli) PullImage(ctx context.Context, image string) error {
	pullOut, err := c.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	defer func() {
		if pullOut != nil {
			pullOut.Close()
		}
	}()
	if err != nil {
		return err
	}

	var e errorMessage
	buffIOReader := bufio.NewReader(pullOut)
	for {
		streamBytes, err := buffIOReader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if c.log != nil {
			_, _ = c.log.Write(streamBytes)
		}

		_ = json.Unmarshal(streamBytes, &e)
		if e.Error != "" {
			return errors.New(e.Error)
		}
	}

	return nil
}

func (c *Cli) PushImage(ctx context.Context, image string) error {
	pushOut, err := c.cli.ImagePush(ctx, image, types.ImagePushOptions{
		RegistryAuth: c.auth,
	})
	defer func() {
		if pushOut != nil {
			pushOut.Close()
		}
	}()
	if err != nil {
		return err
	}

	var e errorMessage
	buffIOReader := bufio.NewReader(pushOut)
	for {
		streamBytes, err := buffIOReader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if c.log != nil {
			_, _ = c.log.Write(streamBytes)
		}

		_ = json.Unmarshal(streamBytes, &e)
		if e.Error != "" {
			return errors.New(e.Error)
		}
	}

	return nil
}
