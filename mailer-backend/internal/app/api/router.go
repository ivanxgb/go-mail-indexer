package api

import (
	"github.com/go-chi/chi/v5"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)
	setHandlers(router)

	return router
}