package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	database "gateway/database/postgres"
	auth "gateway/middleware/auth/jwt"
)

var (
	serviceDB database.IDatabase
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
	dbConn := database.PGConnProps{
		DB_HOST:     "db",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "user_info",
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	serviceDB = database.Initial(dbConn)
	db, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
	} else {
		serviceAuth := auth.ServiceAuth(db)
		// Server header
		e.Use(ServerHeader)

		// Routes
		e.POST("/signup", Signup)
		e.POST("/login", serviceAuth.Login)
		e.GET("/health", healthcheck)

		secGroup := e.Group("/api")
		{
			secGroup.Use(serviceAuth.IsAuth)
			secGroup.GET("/v1/votes", restricted)
		}
	}

	e.Logger.Fatal(e.Start(":1323"))
}
