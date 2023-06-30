package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := c.Request()
		headers := req.Header

		bearerToken := headers.Get("Authorization")
		if !strings.HasPrefix(bearerToken, "Bearer") {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

		fmt.Printf("tokenString: %s\n", tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		}, jwt.WithLeeway(5*time.Second))

		if token.Valid {
			claims, ok := token.Claims.(*JwtCustomClaims)
			if ok {
				fmt.Printf("Name: %v\n", claims.Name)
				fmt.Printf("Admin: %v\n", claims.Admin)
				fmt.Printf("Issuer: %v\n", claims.Issuer)
			} else {
				fmt.Printf("Is Claims: %v\n", ok)
			}

		} else {
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

			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
