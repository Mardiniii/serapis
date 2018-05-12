package evaluator

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func pullImage(cli *client.Client, imgName string) (io.Reader, error) {
	reader, err := cli.ImagePull(ctx, imgName, types.ImagePullOptions{})
	return reader, err
}

func logContainer(cli *client.Client, id string) (io.Reader, error) {
	output, err := cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})

	return output, err
}

func createContainer(cli *client.Client, img string, cmd []string) (container.ContainerCreateCreatedBody, error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Tty:       true,
		OpenStdin: true,
		StdinOnce: true,
		Image:     img,
		Cmd:       cmd,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/Users/sebastianzapatamardini/go/src/github.com/Mardiniii/serapis/tmp/scripts",
				Target: "/scripts",
			},
		},
	}, nil, "")

	return resp, err
}

func startContainer(cli *client.Client, id string) error {
	err := cli.ContainerStart(ctx, id, types.ContainerStartOptions{})
	return err
}

func attachContainer(cli *client.Client, id string, stdin []string) error {
	log.Println("Received Stdin in command")
	opts := types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
	}
	resp, err := cli.ContainerAttach(context.Background(), id, opts)
	if err != nil {
		return err
	}
	defer resp.Close()

	str := strings.Join(stdin, "\n") + "\n"
	log.Println("Writing to Stdin: " + str)
	resp.Conn.Write([]byte(str))

	return nil
}

func waitContainer(cli *client.Client, id string) int {
	statusCh, errCh := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
		return 1
	case okBody := <-statusCh:
		return int(okBody.StatusCode)
	}
}

func removeContainer(cli *client.Client, id string) error {
	err := cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	})

	return err
}
