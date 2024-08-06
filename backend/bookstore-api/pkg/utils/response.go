package utils

import "github.com/labstack/echo/v4"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func SendResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
