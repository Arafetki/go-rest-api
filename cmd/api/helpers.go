package main

import "github.com/golang-jwt/jwt/v5"

type envelope map[string]any

type Claims struct {
	Email string
	jwt.RegisteredClaims
}
