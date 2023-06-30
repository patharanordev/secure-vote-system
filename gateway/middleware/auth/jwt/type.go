package auth

import "github.com/golang-jwt/jwt/v5"

type UserAccount struct {
	ID    string
	Name  string
	Admin bool
}

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
