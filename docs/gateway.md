# **API Gateway**

This service provided middlewares to handle any incoming traffic before accessing to our microservices.

## **Middlewares**

The middleware including :

- JWT Authorization service — users **MUST** authorized before get info from our APIs.
- Attached `x-user-id` or user ID into request header before calling our microservices. The ID will be referenced both frontend/backend :

    ```go
    // Ref. `gateway/middleware/auth/jwt/jwt.go`
    c.Request().Header.Set(s.ReqHeader.UserID, account.ID)
    ```

- HTTP Request — after user authorized, the service will calling target endpoint based on client required.

    ***Example*** do healthcheck our microservice :

    *Setup route*

    ```go
    secGroup.GET("/v1/votes", restricted)
    ```

    *HTTP request*

    ```go
    func restricted(c echo.Context) error {
        resp, errRes := http.Get("http://apis:1323/health")
        if errRes != nil {
            return handleError(c, errRes)
        }

        body, errReadBody := ioutil.ReadAll(resp.Body)
        if errReadBody != nil {
            handleError(c, errReadBody)
        }

        //Convert the body to type string
        sb := string(body)
        fmt.Printf("Calling internal API: %v\n", sb)

        return c.JSON(http.StatusOK, &res.ResponseObject{
            Status: http.StatusOK,
            Data:   &sb,
            Error:  nil,
        })
    }
    ```

    *Response object*

    ```json
    {
        "status": 200,
        "data": "OK",
        "error": null
    }
    ```

- Declare response schema object pattern.

    ```go
    type ResponseObject struct {
        Status int         `json:"status"`
        Data   interface{} `json:"data"`
        Error  *string     `json:"error"`
    }
    ```

- Logging pattern.

## **Services**

Generally I will separate account management service (sign-in/out/up) out to another service. In this situation, I attached the service into API gateway.

> **CRUD**
>
> Ref. `gateway/middleware/auth/jwt/account.go`.

- Login
- Signup
- UpdateAccount
- DeleteAccount

This service requires `database` module, to calling any CRUD function, you **MUST** executing inside `Connect()`/`Close()` block :

```go
// ...

_, errDB := serviceDB.Connect()

if errDB != nil {
    fmt.Printf("Connect to database error : %s\n", errDB.Error())
    return errDB
}

/* --- Execute your CRUD function --- */

serviceDB.Close()

// Do something next...
```
