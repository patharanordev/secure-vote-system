package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "gateway/middleware/auth/jwt"
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
	return c.String(http.StatusOK, "Vote list")
}

func main() {
	e := echo.New()
	serviceAuth := auth.ServiceAuth()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Server header
	e.Use(ServerHeader)

	// Routes
	e.POST("/login", serviceAuth.Login)
	e.GET("/health", healthcheck)

	secGroup := e.Group("/api")
	{
		secGroup.Use(serviceAuth.IsAuth)
		secGroup.GET("/v1/votes", restricted)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
