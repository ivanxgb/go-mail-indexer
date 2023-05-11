package init_app

import (
	de "indexer/dir_explorer"
	"indexer/env_loader"
	"indexer/utils"
	zu "indexer/zinc_uploader"
	"os"
)

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
	if len(os.Args) < 2 {
		utils.ErrorPrinter("No directory path provided")
	}

	dirPath := os.Args[1]
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
