package evaluator

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Mardiniii/serapis/common/models"
	"github.com/docker/docker/client"
)

var ctx = context.Background()

var images = map[string]string{
	"node": "node:latest",
	"ruby": "ruby:latest",
}

var extensions = map[string]string{
	"node": "js",
	"ruby": "rb",
}

func checkError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func copyLogsToStdOut(output io.Reader) {
	io.Copy(os.Stdout, output)
}

func parseLogsToString(output io.Reader) (string, error) {
	b, err := ioutil.ReadAll(output)
	if err != nil {
		return "", err
	}
	return string(b), nil

}

// Start uses the params givent to run a piece code into an isolated container
func Start(eval *models.Evaluation) {
	var err error
	img := images[eval.Language]

	// Create a temporary file with the code to evaluate
	fileName, err := createFile(eval.Language, eval.Code)
	checkError(err)

	cmd := []string{eval.Language, "/scripts/" + fileName}

	cli, err := client.NewEnvClient()
	checkError(err)

	// Pull image
	reader, err := pullImage(cli, img)
	checkError(err)
	copyLogsToStdOut(reader)

	// Create container
	resp, err := createContainer(cli, img, cmd)
	checkError(err)

	// Start container
	err = startContainer(cli, resp.ID)
	checkError(err)

	// Attach to container is eval.Stdin has values
	if len(eval.Stdin) > 0 {
		err = attachContainer(cli, resp.ID, eval.Stdin)
		checkError(err)
	}

	// Wait for container
	eval.ExitCode = waitContainer(cli, resp.ID)

	// Log container in the output Reader
	output, err := logContainer(cli, resp.ID)
	checkError(err)

	// Parse logs as string to be returned
	containerOuput, err := parseLogsToString(output)
	checkError(err)
	eval.Output = containerOuput

	// Remove container before exit
	err = removeContainer(cli, resp.ID)
	checkError(err)

	// Remove the file from the /tmp/scripts directory after finishing
	err = removeFile(fileName)
	checkError(err)
}
