package app

import (
	"flag"
	"indexer/env_loader"
	de "indexer/internal/app/dir_explorer"
	zu "indexer/internal/app/services/zinc_uploader"
	"indexer/internal/app/utils"
	"os"
)

var mailsDirPath = flag.String("mails", "/Users/ivanxgb/Developer/mails", "Path to the directory that contains the emails to be indexed")

func Init() {
	loadEnv()
	dirPath := getDirectoryPath()
	filePaths, err := de.GetFilePaths(dirPath)

	if err != nil {
		utils.ErrorPrinter("There was an error getting the paths of the files")
	}

	zu.ProcessFiles(&filePaths)
}

func getDirectoryPath() string {
	flag.Parse()

	dirPath := *mailsDirPath

	if dirPath == "" {
		utils.ErrorPrinter("No directory path provided")
	}

	checkDirectory(dirPath)

	return dirPath
}

func checkDirectory(dirPath string) {
	dirInfo, err := os.Stat(dirPath)

	if err != nil || !dirInfo.IsDir() {
		utils.ErrorPrinter("Invalid directory path provided")
	}
}

func loadEnv() {
	env_loader.ExportEnv()
}
