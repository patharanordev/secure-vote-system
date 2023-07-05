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

func HTTPRequest(c echo.Context, url string) error {
	bodyBytes, errBody := ioutil.ReadAll(c.Request().Body)
	if errBody != nil {
		return res.HandleError(c, http.StatusBadRequest, errBody)
	}

	req, errReq := http.NewRequest(c.Request().Method, url, bytes.NewBuffer(bodyBytes))
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
