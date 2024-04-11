package main

import (
	"expvar"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Use(middleware.Heartbeat("/ping"))

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.checkHealthHandler)
		r.Get("/metrics", expvar.Handler().ServeHTTP)

	})

	return router
}
