package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetVoteItem(c echo.Context) error {
	id := c.QueryParam("id")
	url := fmt.Sprintf("http://apis:1323/vote-item?id=%s", id)
	return HTTPRequest(c, url)
}

func AddVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}

func UpdateVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}

func Voting(c echo.Context) error {
	url := "http://apis:1323/voting"
	return HTTPRequest(c, url)
}

func ClearVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}
