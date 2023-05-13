package api

import (
	"github.com/go-chi/chi/v5"
	"mailer-backend/internal/app/handler"
)

const (
	searchRoute = "/search"
)

func setHandlers(router *chi.Mux) {
	router.HandleFunc("/", handler.HomeHandler)
	router.Handle("/assets/*", handler.AssetHandler())
	router.Post(searchRoute, handler.SearchHandler)
}
