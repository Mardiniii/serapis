package evaluator

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Mardiniii/serapis/common/models"
	"github.com/docker/docker/client"
)

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

type executor struct {
	eval *models.Evaluation
}

// Start uses the params givent to run a piece code into an isolated container
func (e *executor) Start() (exitCode int, containerOutput string, err error) {
	var cmd []string
	var codeFileName string
	var runFileName string
	img := images[e.eval.Language]

	if e.eval.Code != "" {
		codeFileName, err = createCodeFile(e.eval.Language, e.eval.Code)
		defer removeFile(codeFileName)
		checkError(err)
	}
	cmd = []string{e.eval.Language, "scripts/" + codeFileName}

	// Create a temporary file with the code to evaluate if a Git repo isn't present
	if string(e.eval.Dependencies) != "null" || string(e.eval.Git) != "null" {
		// Create a temporary .sh file with the commands to install dependencies and
		// run the code file
		runFileName, err = createRunFile(e.eval, codeFileName)
		checkError(err)
		defer removeFile(runFileName)

		cmd = []string{"./scripts/" + runFileName}
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

	// Attach to container is e.eval.Stdin has values
	if len(e.eval.Stdin) > 0 {
		err = attachContainer(cli, resp.ID, e.eval.Stdin)
		checkError(err)
	}

	// Wait for container
	exitCode = waitContainer(cli, resp.ID)

	// Log container in the output Reader
	output, err := logContainer(cli, resp.ID)
	checkError(err)

	// Parse logs as string to be returned
	containerOutput, err = parseLogsToString(output)
	checkError(err)

	// Remove container before exit
	err = removeContainer(cli, resp.ID)
	checkError(err)

	return
}
