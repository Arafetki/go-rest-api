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

	router.Use(middleware.Heartbeat("/ping"))

	router.Route("/v1", func(r chi.Router) {
		// Swagger
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%d/v1/swagger/doc.json", app.cfg.httpPort)),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		),
		)
		r.Get("/healthcheck", app.checkHealthHandler)
		r.Get("/metrics", expvar.Handler().ServeHTTP)

	})

	return router
}
