package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ServiceAuth(db *sql.DB) IServiceAuth {
	// Initial
	s := &ServiceAuthProps{}
	s.ReqHeader.Authorization = "Authorization"
	s.ReqHeader.UserID = "x-user-id"
	s.JwtSecret = "secret"
	s.AuthType = "Bearer"
	s.ResMsg.Unauthorized = "Unauthorized"
	s.DB = db

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

func (s *ServiceAuthProps) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userAccount, errAccount := s.getUser(username, password)
	// Throws unauthorized error
	if errAccount != nil {
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

func (s *ServiceAuthProps) getUser(username string, password string) (*UserAccount, error) {

	// TODO: Get user from storage

	if username != "jon" || password != "shhh!" {
		return nil, errors.New(s.ResMsg.Unauthorized)
	}

	return &UserAccount{
		"07cc32b5-4e73-4a2e-bab5-8de399a41df5",
		"Jon Snow",
		true,
	}, nil
}
