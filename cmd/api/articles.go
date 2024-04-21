package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/Arafetki/my-portfolio-api/internal/request"
	"github.com/Arafetki/my-portfolio-api/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {

	user := app.contextGetUser(r)

	var input struct {
		Title         string `json:"title"`
		Body          string `json:"body"`
		CategoriesIds []int  `json:"categories_ids"`
		Published     bool   `json:"published"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	article := &models.Article{
		Title:       input.Title,
		Body:        input.Body,
		Published:   input.Published,
		PublishDate: "0001-01-01",
	}
	article.Author = user.Email
	if article.Published {
		article.PublishDate = time.Now().Format("2006-01-02")
	}

	err = app.validator.Struct(article)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		app.failedValidationResponse(w, r, validationErrors)
		return
	}

	err = app.repository.Article.Create(article)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, article)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}

}

func (app *application) fetchArticleHandler(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParamFromCtx(r.Context(), "id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, errors.New("invalid id param"))
		return
	}

	article, err := app.repository.Article.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
		return
	}
	err = response.JSON(w, http.StatusOK, envelope{"article": article})
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
