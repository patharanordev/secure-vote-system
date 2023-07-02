package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	database "gateway/database/postgres"
	res "gateway/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	serviceDB database.IDatabase
)

func ServiceAuth() IServiceAuth {
	// Initial
	s := &ServiceAuthProps{}
	s.ReqHeader.Authorization = "Authorization"
	s.ReqHeader.UserID = "x-user-id"
	s.JwtSecret = "secret"
	s.AuthType = "Bearer"
	s.ResMsg.Unauthorized = "Unauthorized"

	dbConn := database.PGConnProps{
		DB_HOST:     "db",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "user_info",
	}

	serviceDB = database.Initial(dbConn)
	return s
}

func (s *ServiceAuthProps) IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := c.Request()
		headers := req.Header
		AuthTypePrefix := fmt.Sprintf("%s ", s.AuthType)

		bearerToken := headers.Get(s.ReqHeader.Authorization)
		if !strings.HasPrefix(bearerToken, s.AuthType) {
			return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
				Status: http.StatusUnauthorized,
				Data:   nil,
				Error:  &s.ResMsg.Unauthorized,
			})
		}
		tokenString := strings.TrimPrefix(bearerToken, AuthTypePrefix)

		fmt.Printf("tokenString: %s\n", tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.JwtSecret), nil
		}, jwt.WithLeeway(5*time.Second))

		if !token.Valid {
			errMsg := ""
			// Just logging to our system...
			if errors.Is(err, jwt.ErrTokenMalformed) {
				errMsg = fmt.Sprint("That's not even a token")
			} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				// Invalid signature
				errMsg = fmt.Sprint("Invalid signature")
			} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				// Token is either expired or not active yet
				errMsg = fmt.Sprint("Timing is everything")
			} else {
				errMsg = fmt.Sprint("Couldn't handle this token:", err)
			}

			fmt.Printf("Token invalid : %s\n", errMsg)

			return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
				Status: http.StatusUnauthorized,
				Data:   nil,
				Error:  &errMsg,
			})
		}

		claims, ok := token.Claims.(*JwtCustomClaims)
		if !ok {
			errMsg := fmt.Sprint("Cannot claims")
			return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
				Status: http.StatusUnauthorized,
				Data:   nil,
				Error:  &errMsg,
			})
		}

		fmt.Printf(" - ID: %v\n", claims.ID)
		fmt.Printf(" - Name: %v\n", claims.Name)
		fmt.Printf(" - Admin: %v\n", claims.Admin)
		fmt.Printf(" - Issuer: %v\n", claims.Issuer)

		account, errAccount := s.getUserByID(claims.ID)

		if errAccount != nil {
			errMsg := errAccount.Error()
			return c.JSON(http.StatusUnauthorized, &res.ResponseObject{
				Status: http.StatusUnauthorized,
				Data:   nil,
				Error:  &errMsg,
			})
		}

		c.Response().Header().Set(s.ReqHeader.UserID, account.ID)

		return next(c)
	}
}

func (s *ServiceAuthProps) getUser(username string, password string) (*UserAccount, error) {

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return nil, errDB
	}

	account, errAccount := serviceDB.GetAccount(username, password)
	serviceDB.Close()

	if errAccount != nil {
		return nil, errAccount
	}

	return &UserAccount{
		string(account.UID),
		account.Info.Username,
		account.Info.IsAdmin,
	}, nil
}

func (s *ServiceAuthProps) getUserByID(uid string) (*UserAccount, error) {

	_, errDB := serviceDB.Connect()

	if errDB != nil {
		fmt.Printf("Connect to database error : %s\n", errDB.Error())
		return nil, errDB
	}

	account, errAccount := serviceDB.GetAccountByID(uid)
	serviceDB.Close()

	if errAccount != nil {
		return nil, errAccount
	}

	return &UserAccount{
		string(account.UID),
		account.Info.Username,
		account.Info.IsAdmin,
	}, nil
}

func (s *ServiceAuthProps) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userAccount, errAccount := s.getUser(username, password)
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
