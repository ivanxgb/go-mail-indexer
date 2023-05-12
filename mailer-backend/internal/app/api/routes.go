package api

import (
	"github.com/go-chi/chi/v5"
	"mailer-backend/internal/app/handler"
)

const (
	searchRoute = "/search"
)

func setRoutes(router *chi.Mux) {
	router.Post(searchRoute, handler.SearchHandler)
}
