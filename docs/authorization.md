# **Authentication/Authorization**

> Ref. `gateway/middleware/auth/jwt/jwt.go`

## **Authentication**

Wrap your private/internal APIs with specific route. In this case, it is `/api`. Under `/api` will calling middleware name `IsAuth()` :

```go
serviceAuth = auth.ServiceAuth()
secGroup := e.Group("/api")
{
    secGroup.Use(serviceAuth.IsAuth)

    // --- Adding your private endpoint ---
    //
    // Ex. 
    // secGroup.GET("/v1/votes", restricted)
}
```

If the token is valid or user account authorized, client can access to our resources.

## **Authorization**

Firstly, it will check token invalid or not.

```go
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
}
```

Then claim the token to get info :

```go
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
```

But it's not guaruntee that info is in our system, please check it in your database too :

```go
account, errAccount := s.getUserByID(claims.ID)
```

In case it valided, don't for get attached user ID into request header to be a reference in your service.

```go
c.Request().Header.Set(s.ReqHeader.UserID, account.ID)
```

For any failure, please point to error code `401` only to prevent hacker understand your mechanism.

## **Unauthorized**

Client will be received `401` if the account unauthorized, it can be :

- Account not found in the system.
- Username/Password are not match.
- Wrong token.
- Wrong schema.
- ...
