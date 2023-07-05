package main

import (
	"fmt"
	"net/http"
	"strings"

	database "apis/database/postgres"
	res "apis/response"

	"github.com/labstack/echo/v4"
)

func CreateVoteItem(c echo.Context) error {
	payload := new(database.CreateVoteItemPayload)
	if err := c.Bind(payload); err != nil {
		fmt.Printf(" - CreateVoteItem error: %v\n", err)
		errMsg := "Your payload should contains 'itemName', and 'itemDescription'."
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	fmt.Printf(" - CreateVoteItem Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	lastInsertId, errExec := serviceDB.CreateVoteItem(payload)
	serviceDB.Close()

	if errExec != nil {
		errExecMsg := errExec.Error()
		reason := errExecMsg

		fmt.Printf("Insert to database error : %s", errExecMsg)
		if strings.Contains(errExecMsg, "duplicate key") {
			reason = "The user name already exists."
		} else {
			reason = "Cannot create the account."
		}

		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &reason,
		})
	}

	fmt.Println("Created, last inserted id : ", lastInsertId)

	return c.JSON(http.StatusCreated, &res.ResponseObject{
		Status: http.StatusCreated,
		Data:   "Vote item created.",
		Error:  nil,
	})
}

func UpdateVoteItemByID(c echo.Context) error {
	payload := new(database.VoteItemPayload)
	if err := c.Bind(payload); err != nil {
		fmt.Printf("UpdateVoteItemByID error : %v\n", err.Error())
		errMsg := "Your payload should contains 'id'."
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	fmt.Printf("Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	errExec := serviceDB.UpdateVoteItemByID(payload)
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
		Data:   "Vote item updated.",
		Error:  nil,
	})
}

func DeleteVoteItemByID(c echo.Context) error {
	payload := new(database.VoteItemIDPayload)
	if err := c.Bind(payload); err != nil {
		errMsg := "Your payload should contains 'id'."
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	fmt.Printf("Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	errExec := serviceDB.DeleteVoteItemByID(payload)
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
		Data:   "Vote item deleted.",
		Error:  nil,
	})
}
