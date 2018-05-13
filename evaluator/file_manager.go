package evaluator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mardiniii/serapis/common/models"
)

func createCodeFile(lang, code string) (string, error) {
	fileName := lang + "." + extensions[lang]
	filePath, _ := filepath.Abs("../serapis/tmp/scripts/" + fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fmt.Fprintf(file, code)
	return fileName, nil
}

func createRunFile(eval *models.Evaluation, codeFileName string) (string, error) {
	manager := packageManagers[eval.Language]
	fileName := "solution.sh"
	filePath, _ := filepath.Abs("../serapis/tmp/scripts/" + fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Update file permissions before copying in container
	err = os.Chmod(filePath, 0777)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(file, "#!/bin/bash")

	if eval.Git["repo"] != "" {
		fmt.Fprintln(file, "git clone "+eval.Git["repo"]+".git gitrepo")
		fmt.Fprintln(file, "cd gitrepo/")
	}

	for dependency, version := range eval.Dependencies {
		var target string

		if version == "latest" {
			target = dependency
		} else {
			target = dependency + manager["versioner"] + version
		}

		fmt.Fprintln(file, manager["installer"]+target)
	}

	if codeFileName != "" {
		fmt.Fprintln(file, eval.Language+" /scripts/"+codeFileName)
	} else {
		fmt.Fprintln(file, eval.Git["command"])
	}

	return fileName, nil
}

func removeFile(fileName string) error {
	filePath, err := filepath.Abs("../serapis/tmp/scripts/" + fileName)
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
