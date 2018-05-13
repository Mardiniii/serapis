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

var packageManagers = map[string]map[string]string{
	"node": map[string]string{
		"installer": "npm install ",
		"versioner": "@",
	},
	"ruby": map[string]string{
		"installer": "gem install ",
		"versioner": ":",
	},
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
	var cmd []string
	var codeFileName string
	var runFileName string
	img := images[eval.Language]

	// Create a temporary file with the code to evaluate if a Git repo isn't present
	if len(eval.Git) == 0 {
		codeFileName, err = createCodeFile(eval.Language, eval.Code)
		checkError(err)
		defer removeFile(codeFileName)
	}

	// Dependencies installation
	if len(eval.Dependencies) > 0 || len(eval.Git) > 0 {
		// Create a temporary .sh file with the commands to install dependencies and
		// run the code file
		runFileName, err = createRunFile(eval, codeFileName)
		checkError(err)
		defer removeFile(runFileName)

		cmd = []string{"./scripts/" + runFileName}
	} else if len(eval.Stdin) > 0 {
		cmd = []string{eval.Language, "/scripts/" + codeFileName}
	}

	// Create a Docker client
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
}
