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
		fmt.Println(path)
		fmt.Println("Files found: ", filesCount)

		return nil
	})

	if err != nil {
		utils.ErrorPrinter("There was an error reading the directory")
	}
}
