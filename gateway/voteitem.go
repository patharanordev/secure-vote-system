package main

import (
	"github.com/labstack/echo/v4"
)

func AddVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}

func UpdateVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}

func ClearVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	return HTTPRequest(c, url)
}
