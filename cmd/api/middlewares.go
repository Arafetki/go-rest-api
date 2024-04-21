package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Arafetki/my-portfolio-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
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

		t, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(app.cfg.jwt.secretkey), nil
		})

		if err != nil {
			app.logger.Error(err.Error())
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		if claims, ok := t.Claims.(*Claims); ok && t.Valid {
			r = app.contextSetUser(r, &models.User{
				ID:    claims.Subject,
				Email: claims.Email,
			})
		} else {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)

	})
}
