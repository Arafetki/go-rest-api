package main

import (
	"net/http"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/Arafetki/my-portfolio-api/internal/request"
	"github.com/Arafetki/my-portfolio-api/internal/response"
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
