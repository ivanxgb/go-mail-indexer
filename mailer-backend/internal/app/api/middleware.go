package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// setMiddlewares sets the middlewares for the router.
func setMiddlewares(router *chi.Mux) {
	router.Use(
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)
}
