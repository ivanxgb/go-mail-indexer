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
	filesPath, err := de.GetFilesPath(dirPath)

	if err != nil {
		utils.ErrorPrinter("There was an error getting the files path")
	}

	zu.SendFilesToServer(filesPath)
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
