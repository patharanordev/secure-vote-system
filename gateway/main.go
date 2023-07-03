package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "gateway/middleware/auth/jwt"
	res "gateway/response"
)

var (
	serviceAuth auth.IServiceAuth
)

// Custom header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "PatharaNor")
		return next(c)
	}
}

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func restricted(c echo.Context) error {
	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   &res.VotingInfos{Votes: []string{"a"}},
		Error:  nil,
	})
}

func updateAccount(c echo.Context) error {
	return serviceAuth.UpdateAccount(c)
}

func deleteAccount(c echo.Context) error {
	return serviceAuth.DeleteAccount(c)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	serviceAuth = auth.ServiceAuth()
	// Server header
	e.Use(ServerHeader)

	// Routes
	e.POST("/signup", serviceAuth.Signup)
	e.POST("/login", serviceAuth.Login)
	e.GET("/health", healthcheck)

	secGroup := e.Group("/api")
	{
		secGroup.Use(serviceAuth.IsAuth)
		secGroup.GET("/v1/votes", restricted)
		secGroup.POST("/v1/account", updateAccount)
		secGroup.DELETE("/v1/account", deleteAccount)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
