package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	database "apis/database/postgres"
)

var (
	serviceDB database.IDatabase
)

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()
	fmt.Printf("Current timestamp : %v", time.Now())

	dbConn := database.PGConnProps{
		DB_HOST:     "db",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "user_info",
	}

	serviceDB = database.Initial(dbConn)

	// ----- Middleware -----
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ----- Routes -----
	e.GET("/health", healthcheck)

	// Vote item
	e.POST("/vote-item", CreateVoteItem)
	e.PATCH("/vote-item", UpdateVoteItemByID)
	e.DELETE("/vote-item", DeleteVoteItemByID)

	// Vote list
	e.GET("/votes", GetVoteList)
	e.DELETE("/votes", DeleteVoteList)

	// Voting
	e.PATCH("/voting", Voting)

	e.Logger.Fatal(e.Start(":1323"))
}
