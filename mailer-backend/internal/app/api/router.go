package api

import (
	"github.com/go-chi/chi/v5"
)

// GetRouter returns a new chi router with the middlewares and handlers set.
func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)
	setHandlers(router)

	return router
}
