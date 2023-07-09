package main

import (
	"fmt"
	database "gateway/database/postgres"
	res "gateway/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVotes(c echo.Context) error {
	url := "http://apis:1323/votes"
	return HTTPRequest(c, url)
}

func ClearVotes(c echo.Context) error {
	url := "http://apis:1323/votes"

	dbConn := database.PGConnProps{
		DB_HOST:     "db",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "user_info",
	}

	serviceDB := database.Initial(dbConn)

	errMsg := "Unauthorized"
	userId := c.Request().Header.Get("x-user-id")
	fmt.Printf(" - User ID : %s\n", userId)
	if len(userId) <= 0 {
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	// Check role/permission
	errMsg = "Your role cannot access to the resource."
	account, errExec := serviceDB.GetAccountByID(userId)
	serviceDB.Close()

	if errExec != nil {
		errMsg = errExec.Error()
		fmt.Printf("Execute SQL error : %s", errMsg)
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	if !account.Info.IsAdmin {
		return c.JSON(http.StatusForbidden, &res.ResponseObject{
			Status: http.StatusForbidden,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	return HTTPRequest(c, url)
}
