package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFile(lang, code string) (string, error) {
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
