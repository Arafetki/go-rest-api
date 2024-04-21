package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/nedpals/supabase-go"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			r = app.contextSetUser(r, models.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Todo : Validate Token
		tokenString := headerParts[1]

		supabaseClient := supabase.CreateClient(app.cfg.supabase.url, app.cfg.supabase.key)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user, err := supabaseClient.Auth.User(ctx, tokenString)
		if err != nil {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}
		r = app.contextSetUser(r, &models.User{
			ID:    user.ID,
			Email: user.Email,
		})

		next.ServeHTTP(w, r)

	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := app.contextGetUser(r)
		if user.IsAnonymous() {
			app.authenticationRequiredResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
