package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
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

func createRunFile(lang, codeFileName string, dependencies map[string]string) (string, error) {
	manager := packageManagers[lang]
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
	for dependency, version := range dependencies {
		var target string

		if version == "latest" {
			target = dependency
		} else {
			target = dependency + manager["versioner"] + version
		}

		fmt.Fprintln(file, manager["installer"]+target)
	}
	fmt.Fprintln(file, lang+" /scripts/"+codeFileName)

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
