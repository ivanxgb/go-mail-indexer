package handler

import (
	"io/fs"
	"mailer-backend/ui"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.Frontend.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := ui.Frontend.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func AssetHandler() http.Handler {
	staticFs, _ := fs.Sub(ui.Frontend, "dist")

	return http.FileServer(http.FS(staticFs))
}
