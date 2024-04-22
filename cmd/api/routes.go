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

	// @General : middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Route("/v1", func(r chi.Router) {

		r.Use(app.authenticate)

		// @API V1 : Public routes
		r.Get("/healthcheck", app.checkHealthHandler)
		// r.Get("/articles/{id}", app.fetchArticleHandler)
		r.Post("/subs", app.createSubHandler)

		// @API V1 : Private routes
		r.Group(func(privateRouter chi.Router) {

			privateRouter.Use(app.requireAuthenticatedUser)

			privateRouter.Get("/swagger/*", httpSwagger.Handler(
				httpSwagger.URL(fmt.Sprintf("http://localhost:%d/v1/swagger/doc.json", app.cfg.httpPort)),
				httpSwagger.DeepLinking(true),
				httpSwagger.DocExpansion("none"),
				httpSwagger.DomID("swagger-ui"),
			),
			)
			privateRouter.Get("/metrics", expvar.Handler().ServeHTTP)

			privateRouter.Post("/articles", app.createArticleHandler)
			privateRouter.Delete("/articles/{id}", app.deleteArticleHandler)

			privateRouter.Get("/subs", app.fetchAllSubsHandler)
			privateRouter.Delete("/subs/{email}", app.deleteSubHandler)
		})

	})

	return router
}
