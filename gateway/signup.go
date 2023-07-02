package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

func Signup(c echo.Context) error {
	payload := new(CreateUserPayload)
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "Your payload should contains 'username', 'password' and 'isAdmin'.")
	}

	fmt.Printf("Received payload : %v\n", payload)

	lastInsertId, errInserted := serviceDB.CreateAccount(payload.Username, payload.Password)

	if errInserted != nil {
		errInsertedMsg := errInserted.Error()
		reason := errInsertedMsg

		fmt.Printf("Insert to database error : %s", errInsertedMsg)
		if strings.Contains(errInsertedMsg, "duplicate key") {
			reason = "The user name already exists."
		} else {
			reason = "Cannot create the account."
		}

		return c.String(http.StatusBadRequest, reason)
	}

	fmt.Println("Created, last inserted id : ", lastInsertId)

	return c.String(http.StatusCreated, "Account created.")
}
