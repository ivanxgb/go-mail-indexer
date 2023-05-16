package app

import (
	"flag"
	"indexer/internal/app/utils"
	"os"
	"path/filepath"
)

const (
	// extensionFile is the extension of the mails files.
	extensionFile = "."
)

var (
	// mailsDirPath is the path to the directory that contains the emails to be indexed.
	mailsDirPath = flag.String("mails", "/Users/ivanxgb/Developer/mails", "Path to the directory that contains the emails to be indexed")

	// zIndex is the name of the index where the emails will be indexed.
	zIndex = flag.String("index", "en_mails", "Name of the index where the emails will be indexed")
)

// getDirectoryPath checks if the directory path provided is valid and returns it.
func getDirectoryPath() string {
	dirPath := *mailsDirPath

	if dirPath == "" {
		utils.ErrorPrinter("No directory path provided")
	}

	checkDirectory(dirPath)

	return dirPath
}

// checkDirectory checks if the directory path is valid (exists and is a directory).
func checkDirectory(dirPath string) {
	dirInfo, err := os.Stat(dirPath)

	if err != nil || !dirInfo.IsDir() {
		utils.ErrorPrinter("Invalid directory path provided")
	}
}

// getFilePaths receives a directory path and returns a slice of strings with
// the path of all the files inside the directory.
func getFilePaths(dirPath string) []string {
	var filePaths []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If the path is not a directory or the file extension is valid, is added to the slice.
		if !info.IsDir() && filepath.Ext(path) == extensionFile {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		utils.ErrorPrinter("There was an error getting the paths of the files")
	}

	return filePaths
}
