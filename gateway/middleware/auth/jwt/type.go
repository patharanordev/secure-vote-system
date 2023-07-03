package auth

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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

type RequestHeaderProps struct {
	Authorization string
	UserID        string
}

type ResponseMessageProps struct {
	Unauthorized string
}

type ServiceAuthProps struct {
	JwtSecret string
	AuthType  string
	ReqHeader RequestHeaderProps
	ResMsg    ResponseMessageProps
	DB        *sql.DB
}

type IServiceAuth interface {
	IsAuth(next echo.HandlerFunc) echo.HandlerFunc
	Login(c echo.Context) error
	Signup(c echo.Context) error
	UpdateAccount(c echo.Context) error
	DeleteAccount(c echo.Context) error
}

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UpdateAccountPayload struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}
