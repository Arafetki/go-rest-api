package main

import (
	"errors"
	"net/http"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/Arafetki/my-portfolio-api/internal/repository"
	"github.com/Arafetki/my-portfolio-api/internal/request"
	"github.com/Arafetki/my-portfolio-api/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (app *application) createSubHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email string `json:"email"`
	}

	err := request.DecodeJSONStrict(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sub := &models.Subscriber{
		Email: input.Email,
	}

	err = app.validator.Struct(sub)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		app.failedValidationResponse(w, r, validationErrors)
		return
	}

	err = app.repository.Subscriber.Create(sub)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateKeyValue):
			app.errorMessage(w, r, http.StatusConflict, "The resource already exists and could not be created", nil)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}
	err = response.JSON(w, http.StatusCreated, envelope{"message": "Subscriber successfully added"})

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}

}

func (app *application) fetchAllSubsHandler(w http.ResponseWriter, r *http.Request) {

	subs, err := app.repository.Subscriber.GetAllEmails()
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, envelope{"subs": subs})
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) deleteSubHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email string `validate:"required,email"`
	}

	input.Email = chi.URLParamFromCtx(r.Context(), "email")

	err := app.validator.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		app.failedValidationResponse(w, r, validationErrors)
		return
	}

	err = app.repository.Subscriber.Delete(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}

	err = response.JSON(w, http.StatusOK, envelope{"message": "Subscriber successfully deleted"})

	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}

}
