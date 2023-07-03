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

		c.Request().Header.Set(s.ReqHeader.UserID, account.ID)

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
