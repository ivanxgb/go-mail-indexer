package handler

import (
	"mailer-backend/ui"
	"net/http"
)

// HomeHandler handles the home route and serves the index.html file from the
// ui package (frontend).
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

// AssetHandler handles the assets route and serves the static files from the
// ui package (frontend).
func AssetHandler() http.Handler {
	staticFs, _ := ui.GetAssets()
	return http.FileServer(http.FS(staticFs))
}
