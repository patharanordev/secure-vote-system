package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()
	fmt.Printf("Current timestamp : %v", time.Now())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", healthcheck)

	e.Logger.Fatal(e.Start(":1323"))
}
