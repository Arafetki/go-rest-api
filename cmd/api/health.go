package main

import (
	"net/http"

	"github.com/Arafetki/my-portfolio-api/internal/response"
)

// Todo

func (app *application) checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.cfg.env,
			"version":     version,
		},
	}
	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
