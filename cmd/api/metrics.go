package main

import (
	"net/http"

	"github.com/Arafetki/my-portfolio-api/internal/response"
)

// Todo

func (app *application) reportMetricsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.cfg.env,
			"version":     "1.0.0",
		},
	}
	response.JSON(w, http.StatusOK, data)
}
