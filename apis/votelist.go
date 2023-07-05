package main

import (
	"fmt"
	"net/http"

	res "apis/response"

	"github.com/labstack/echo/v4"
)

func GetVoteList(c echo.Context) error {
	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	voteItems, errExec := serviceDB.GetVoteList()
	serviceDB.Close()

	if errExec != nil {
		errExecMsg := errExec.Error()
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errExecMsg,
		})
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   voteItems,
		Error:  nil,
	})
}

func DeleteVoteList(c echo.Context) error {
	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	errExec := serviceDB.DeleteVoteList()
	serviceDB.Close()

	if errExec != nil {
		errExecMsg := errExec.Error()
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errExecMsg,
		})
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Vote list is cleared.",
		Error:  nil,
	})
}
