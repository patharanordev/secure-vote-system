package main

import (
	"github.com/labstack/echo/v4"
)

func GetVotes(c echo.Context) error {
	url := "http://apis:1323/votes"
	return HTTPRequest(c, url)
}

func ClearVotes(c echo.Context) error {
	url := "http://apis:1323/votes"
	return HTTPRequest(c, url)
}
