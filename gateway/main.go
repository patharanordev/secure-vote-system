package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	resp, errRes := http.Get("http://apis:1323/health")
	if errRes != nil {
		return res.HandleError(c, resp.StatusCode, errRes)
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		res.HandleError(c, http.StatusBadGateway, errReadBody)
	}

	//Convert the body to type string
	sb := string(body)
	fmt.Printf("Calling internal API: %v\n", sb)

	return c.JSON(resp.StatusCode, &res.ResponseObject{
		Status: resp.StatusCode,
		Data:   &sb,
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
	fmt.Printf("Current timestamp : %v", time.Now())

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

		secGroup.POST("/v1/account", updateAccount)
		secGroup.DELETE("/v1/account", deleteAccount)

		// CRUD vote item
		secGroup.POST("/v1/vote-item", AddVoteItem)
		secGroup.PATCH("/v1/vote-item", UpdateVoteItem)
		secGroup.DELETE("/v1/vote-item", ClearVoteItem)

		// Vote list
		secGroup.GET("/v1/votes", GetVotes)
		secGroup.DELETE("/v1/votes", ClearVotes)

		// Voting
		secGroup.PATCH("/v1/voting", Voting)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
