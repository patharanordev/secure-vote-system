package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	database "gateway/database/postgres"

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
			return c.String(http.StatusUnauthorized, s.ResMsg.Unauthorized)
		}
		tokenString := strings.TrimPrefix(bearerToken, AuthTypePrefix)

		fmt.Printf("tokenString: %s\n", tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.JwtSecret), nil
		}, jwt.WithLeeway(5*time.Second))

		if !token.Valid {
			// Just logging to our system...
			if errors.Is(err, jwt.ErrTokenMalformed) {
				fmt.Println("That's not even a token")
			} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				// Invalid signature
				fmt.Println("Invalid signature")
			} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				// Token is either expired or not active yet
				fmt.Println("Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}

			return c.String(http.StatusUnauthorized, s.ResMsg.Unauthorized)
		}

		claims, ok := token.Claims.(*JwtCustomClaims)
		if !ok {
			fmt.Printf("Is Claims: %v\n", ok)
			return c.String(http.StatusUnauthorized, s.ResMsg.Unauthorized)
		}

		fmt.Printf(" - Name: %v\n", claims.Name)
		fmt.Printf(" - Admin: %v\n", claims.Admin)
		fmt.Printf(" - Issuer: %v\n", claims.Issuer)

		c.Response().Header().Set(s.ReqHeader.UserID, claims.ID)

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

func (s *ServiceAuthProps) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userAccount, errAccount := s.getUser(username, password)
	// Throws unauthorized error
	if errAccount != nil {
		fmt.Printf("Login error : %s\n", errAccount.Error())
		return c.String(http.StatusUnauthorized, s.ResMsg.Unauthorized)
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
		return c.String(http.StatusBadRequest, "Your payload should contains 'username', 'password' and 'isAdmin'.")
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

		return c.String(http.StatusBadRequest, reason)
	}

	fmt.Println("Created, last inserted id : ", lastInsertId)

	return c.String(http.StatusCreated, "Account created.")
}
