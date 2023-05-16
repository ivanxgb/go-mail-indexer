package app

import (
	"flag"
	ev "indexer/env_loader"
	"indexer/internal/app/handler"
)

func Init() {
	flag.Parse()
	loadEnv()

	dirPath := getDirectoryPath()
	filePaths := getFilePaths(dirPath)

	handler.ProcessFiles(&filePaths, *zIndex)
}

func loadEnv() {
	ev.ExportEnv()
}
