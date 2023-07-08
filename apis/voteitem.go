package main

import (
	"fmt"
	"net/http"

	database "apis/database/postgres"
	res "apis/response"

	"github.com/labstack/echo/v4"
)

func handleExecError(c echo.Context, errExec error) error {
	errExecMsg := errExec.Error()
	fmt.Printf("Execute SQL error : %s", errExecMsg)
	return c.JSON(http.StatusBadRequest, &res.ResponseObject{
		Status: http.StatusBadRequest,
		Data:   nil,
		Error:  &errExecMsg,
	})
}

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

	errAuth := "Unauthorized"
	userId := c.Request().Header.Get("x-user-id")
	fmt.Printf(" - User ID : %s\n", userId)
	if len(userId) <= 0 {
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errAuth,
		})
	}

	fmt.Printf(" - CreateVoteItem Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	_, errExec := serviceDB.CreateVoteItem(userId, payload)
	serviceDB.Close()

	if errExec != nil {
		return handleExecError(c, errExec)
	}

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

	errAuth := "Unauthorized"
	userId := c.Request().Header.Get("x-user-id")
	fmt.Printf(" - User ID : %s\n", userId)
	if len(userId) <= 0 {
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errAuth,
		})
	}

	fmt.Printf("Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	errExec := serviceDB.UpdateVoteItemByID(userId, payload)
	serviceDB.Close()

	if errExec != nil {
		return handleExecError(c, errExec)
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Vote item updated.",
		Error:  nil,
	})
}

func Voting(c echo.Context) error {
	payload := new(database.VotingPayload)
	if err := c.Bind(payload); err != nil {
		fmt.Printf("Voting error : %v\n", err.Error())
		errMsg := "Your payload should contains 'id' and 'isUp'."
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	errAuth := "Unauthorized"
	userId := c.Request().Header.Get("x-user-id")
	fmt.Printf(" - User ID : %s\n", userId)
	if len(userId) <= 0 {
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errAuth,
		})
	}

	fmt.Printf("Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	var errExec error
	if payload.IsUp {
		errExec = serviceDB.UpVote(userId, payload)
	} else {
		errExec = serviceDB.DownVote(userId, payload)
	}
	serviceDB.Close()

	if errExec != nil {
		return handleExecError(c, errExec)
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Vote success.",
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

	errAuth := "Unauthorized"
	userId := c.Request().Header.Get("x-user-id")
	fmt.Printf(" - User ID : %s\n", userId)
	if len(userId) <= 0 {
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errAuth,
		})
	}

	fmt.Printf("Received payload : %v\n", payload)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	errExec := serviceDB.DeleteVoteItemByID(userId, payload)
	serviceDB.Close()

	if errExec != nil {
		return handleExecError(c, errExec)
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Vote item deleted.",
		Error:  nil,
	})
}
