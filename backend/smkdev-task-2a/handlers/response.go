package handlers

import (
	"github.com/labstack/echo/v4"
)

type TodoResponse struct {
    StatusCode int         `json:"status_code"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data"`
}

func JSONResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, TodoResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
    })
}