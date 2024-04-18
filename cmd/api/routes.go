package main

import (
	"expvar"
	"fmt"
	"net/http"

	_ "github.com/Arafetki/my-portfolio-api/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) routes() http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	// Public
	router.Use(middleware.Heartbeat("/ping"))

	router.Route("/v1", func(r chi.Router) {

		// Authentication middleware
		r.Use(app.authenticate)

		// Public - No bearer token should be provided in headers
		r.Get("/healthcheck", app.checkHealthHandler)

		// Private - Require bearer token to be provided in headers
		// r.Use(app.requireAuthenticatedUser)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%d/v1/swagger/doc.json", app.cfg.httpPort)),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		),
		)
		r.Get("/metrics", expvar.Handler().ServeHTTP)

	})

	return router
}
