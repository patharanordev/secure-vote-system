package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	res "gateway/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// CSRFToken string `json:"csrfToken"`
}

func (s *ServiceAuthProps) Login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")

	payload := new(LoginPayload)
	if err := c.Bind(payload); err != nil {
		errMsg := "Your payload should contains 'username', 'password'."
		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	userAccount, errAccount := s.getUser(payload.Username, payload.Password)
	// Throws unauthorized error
	if errAccount != nil {
		errMsg := errAccount.Error()
		return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
			Status: http.StatusUnauthorized,
			Data:   nil,
			Error:  &errMsg,
		})
	}

	// Set custom claims
	// Ref. https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1
	claims := &JwtCustomClaims{
		userAccount.ID,
		userAccount.Name,
		userAccount.Admin,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "SecVoteSys",
			Subject:   "SecVoteSys_CustomAuth",
			ID:        "1",
			Audience:  []string{"general_user"},
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (s *ServiceAuthProps) Signup(c echo.Context) error {
	payload := new(CreateUserPayload)
	if err := c.Bind(payload); err != nil {
		errMsg := "Your payload should contains 'username', 'password' and 'isAdmin'."
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

	lastInsertId, errInserted := serviceDB.CreateAccount(payload.Username, payload.Password, payload.IsAdmin)
	serviceDB.Close()

	if errInserted != nil {
		errInsertedMsg := errInserted.Error()
		reason := errInsertedMsg

		fmt.Printf("Insert to database error : %s", errInsertedMsg)
		if strings.Contains(errInsertedMsg, "duplicate key") {
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
		Data:   "Account created.",
		Error:  nil,
	})
}

func (s *ServiceAuthProps) UpdateAccount(c echo.Context) error {
	payload := new(UpdateAccountPayload)
	uid := c.Request().Header.Get(s.ReqHeader.UserID)
	fmt.Printf("User ID : %v\n", uid)

	if err := c.Bind(payload); err != nil {
		errMsg := "Your payload should contains 'username' and 'isAdmin'."
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

	errUpdated := serviceDB.UpdateAccount(uid, payload.Username, payload.IsAdmin)
	serviceDB.Close()

	if errUpdated != nil {
		errUpdatedMsg := errUpdated.Error()
		reason := errUpdatedMsg

		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &reason,
		})
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Account updated.",
		Error:  nil,
	})
}

func (s *ServiceAuthProps) DeleteAccount(c echo.Context) error {
	uid := c.Request().Header.Get(s.ReqHeader.UserID)
	fmt.Printf("User ID : %v\n", uid)

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return errDB
	}

	err := serviceDB.DeleteAccountByID(uid)
	serviceDB.Close()

	if err != nil {
		errMsg := err.Error()
		reason := errMsg

		return c.JSON(http.StatusBadRequest, &res.ResponseObject{
			Status: http.StatusBadRequest,
			Data:   nil,
			Error:  &reason,
		})
	}

	return c.JSON(http.StatusOK, &res.ResponseObject{
		Status: http.StatusOK,
		Data:   "Account deleted.",
		Error:  nil,
	})
}
