package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	res "gateway/response"
)

func AddVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	bodyBytes, errBody := ioutil.ReadAll(c.Request().Body)
	if errBody != nil {
		return res.HandleError(c, http.StatusBadRequest, errBody)
	}

	req, errReq := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if errReq != nil {
		return res.HandleError(c, http.StatusBadRequest, errReq)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		fmt.Printf(" - Gateway error making http request: %s\n", errResp)
		return res.HandleError(c, resp.StatusCode, errResp)
	}

	resObj := new(res.ResponseObject)
	json.NewDecoder(resp.Body).Decode(resObj)

	return c.JSON(resp.StatusCode, &resObj)
}

func UpdateVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	bodyBytes, errBody := ioutil.ReadAll(c.Request().Body)
	if errBody != nil {
		return res.HandleError(c, http.StatusBadRequest, errBody)
	}

	req, errReq := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
	if errReq != nil {
		return res.HandleError(c, http.StatusBadRequest, errReq)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		fmt.Printf(" - Gateway error making http request: %s\n", errResp)
		return res.HandleError(c, resp.StatusCode, errResp)
	}

	resObj := new(res.ResponseObject)
	json.NewDecoder(resp.Body).Decode(resObj)

	return c.JSON(resp.StatusCode, &resObj)
}

func ClearVoteItem(c echo.Context) error {
	url := "http://apis:1323/vote-item"
	bodyBytes, errBody := ioutil.ReadAll(c.Request().Body)
	if errBody != nil {
		return res.HandleError(c, http.StatusBadRequest, errBody)
	}

	req, errReq := http.NewRequest("DELETE", url, bytes.NewBuffer(bodyBytes))
	if errReq != nil {
		return res.HandleError(c, http.StatusBadRequest, errReq)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		fmt.Printf(" - Gateway error making http request: %s\n", errResp)
		return res.HandleError(c, resp.StatusCode, errResp)
	}

	resObj := new(res.ResponseObject)
	json.NewDecoder(resp.Body).Decode(resObj)

	return c.JSON(resp.StatusCode, &resObj)
}
