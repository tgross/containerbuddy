package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func run(t *testing.T) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	name := strings.ToLower(t.Name())
	body, err := cli.ContainerCreate(context.Background(),
		&container.Config{
			Env: []string{"env=x", "env2=y"},
			Image: "myimagename",
			Cmd: ["xx", "xxx"],
			Volumes: map[string]struct{}
		},
		&network.NetworkingConfig{
			EndpointsConfig: map[string]*EndpointSettings{

			}
		},
		name,
	)
	if err != nil {
		t.Fatalf("failed to create container: %v", err)
	}
	id := body.ID
	options := types.ContainerStartOptions{}
	err = cli.ContainerStart(
		context.Background(),
		id,
		options,
	)
}

func pokeDocker() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
