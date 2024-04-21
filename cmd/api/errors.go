package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/Arafetki/my-portfolio-api/internal/response"
	"github.com/go-playground/validator/v10"
)

func (app *application) logServerError(r *http.Request, err error) {
	var (
		message string = err.Error()
		method  string = r.Method
		url     string = r.URL.String()
		trace   string = string(debug.Stack())
	)
	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *application) errorMessage(w http.ResponseWriter, r *http.Request, status int, message any, headers http.Header) {
	err := response.JSONWithHeaders(w, status, envelope{"error": message}, headers)
	if err != nil {
		app.logServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	app.logServerError(r, err)
	app.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)

	headers.Add("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	app.errorMessage(w, r, http.StatusUnauthorized, message, headers)
}

func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorMessage(w, r, http.StatusUnauthorized, message, nil)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logServerError(r, err)
	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors validator.ValidationErrors) {
	app.errorMessage(w, r, http.StatusUnprocessableEntity, errors, nil)
}
