package init_app

import (
	"fmt"
	"indexer/utils"
	"os"
)

func Init() {
	dirPath := GetDirectoryPath()
	fmt.Println(dirPath)
}

func GetDirectoryPath() string {
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
