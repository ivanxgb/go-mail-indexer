package handler

import (
	"mailer-backend/ui"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	index, favicon, err := ui.GetStaticFiles()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		w.Write(favicon)
		return
	}

	w.Write(index)
}

func AssetHandler() http.Handler {
	staticFs, _ := ui.GetAssets()
	return http.FileServer(http.FS(staticFs))
}
