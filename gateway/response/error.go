package response

import (
	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, statusCode int, err error) error {
	errMsg := err.Error()
	return c.JSON(statusCode, &ResponseObject{
		Status: statusCode,
		Data:   nil,
		Error:  &errMsg,
	})
}
