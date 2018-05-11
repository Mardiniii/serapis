package evaluator

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

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

// Evaluate uses the params givent to run a piece code into an isolated container
func Evaluate(lang, code string) (int, string) {
	var err error
	img := images[lang]

	// Create a temporary file with the code to evaluate
	fileName, err := createFile(lang, code)
	checkError(err)

	cmd := []string{lang, "/scripts/" + fileName}

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

	// Wait for container
	exitCode := waitContainer(cli, resp.ID)
	fmt.Println(exitCode)

	// Log container in the output Reader
	output, err := logContainer(cli, resp.ID)
	checkError(err)

	// Parse logs as string to be returned
	containerOuput, err := parseLogsToString(output)
	checkError(err)

	// Remove container before exit
	err = removeContainer(cli, resp.ID)
	checkError(err)

	// Remove the file from the /tmp/scripts directory after finishing
	err = removeFile(fileName)
	checkError(err)

	return exitCode, containerOuput
}
