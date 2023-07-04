package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	res "gateway/response"
)

func GetVotes(c echo.Context) error {
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

func ClearVotes(c echo.Context) error {
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
