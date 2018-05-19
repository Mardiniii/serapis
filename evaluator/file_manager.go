package evaluator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mardiniii/serapis/common/models"
)

func createCodeFile(lang, code string) (string, error) {
	fileName := lang + "." + extensions[lang]
	filePath, _ := filepath.Abs("../scripts/" + fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fmt.Fprintf(file, code)
	return fileName, nil
}

func jsonStringToMap(raw json.RawMessage) (m map[string]string, err error) {
	m = make(map[string]string)
	err = json.Unmarshal(raw, &m)

	return
}

func createRunFile(eval *models.Evaluation, codeFileName string) (string, error) {
	manager := packageManagers[eval.Language]
	fileName := "solution.sh"
	filePath, _ := filepath.Abs("../scripts/" + fileName)

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

	// Parse git json.RawMessage
	git, err := jsonStringToMap(eval.Git)
	if err != nil {
		return "", err
	}
	// Parse dependencies json.RawMessage
	dependencies, err := jsonStringToMap(eval.Dependencies)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(file, "#!/bin/bash")
	if git["repo"] != "" {
		fmt.Fprintln(file, "git clone "+git["repo"]+".git gitrepo")
		fmt.Fprintln(file, "cd gitrepo/")
	}

	for dependency, version := range dependencies {
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
		fmt.Fprintln(file, git["command"])
	}

	return fileName, nil
}

func removeFile(fileName string) error {
	filePath, err := filepath.Abs("../scripts/" + fileName)
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
