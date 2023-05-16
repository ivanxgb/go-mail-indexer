package api

import (
	"github.com/go-chi/chi/v5"
	"mailer-backend/internal/app/handler"
)

const (
	searchRoute  = "/api/search"
	summaryRoute = "/api/summary"
)

// setHandlers sets the handlers for the router.
func setHandlers(router *chi.Mux) {
	router.HandleFunc("/", handler.HomeHandler)
	router.Handle("/assets/*", handler.AssetHandler())

	router.Post(searchRoute, handler.SearchHandler)
	router.Post(summaryRoute, handler.SummaryHandler)
}
