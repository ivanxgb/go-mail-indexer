package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var embedFE embed.FS

func GetStaticFiles() (index []byte, favicon []byte, err error) {
	index, err = embedFE.ReadFile("dist/index.html")
	favicon, err = embedFE.ReadFile("dist/favicon.ico")
	return
}

func GetAssets() (fs.FS, error) {
	return fs.Sub(embedFE, "dist")
}
