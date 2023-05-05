package dir_explorer

import (
	"fmt"
	"indexer/utils"
	"os"
	"path/filepath"
)

const (
	extensionFile = "."
)

func GetFilesInDirectory(dirPath string) {
	var filesCount = 0

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories and invalid files
		if info.IsDir() || filepath.Ext(path) != extensionFile {
			return nil
		}

		filesCount++
		fmt.Println("Files found: ", filesCount)
		fmt.Println(readFile(path))
		return nil
	})

	if err != nil {
		utils.ErrorPrinter("There was an error reading the directory")
	}
}

func readFile(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("There was an error opening the file: " + filepath)
		return "", err
	}

	return string(file), nil
}
