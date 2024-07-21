package utils

import (
	"github.com/labstack/echo/v4"
)

// NewErrorResponse membuat response error
// Fungsi ini menerima context Echo, pesan error, dan kode status HTTP, lalu mengembalikan response JSON dengan format yang sudah ditentukan
func NewErrorResponse(c echo.Context, message string, statusCode int) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  statusCode,
	})
}

// NewSuccessResponse membuat response sukses
// Fungsi ini menerima context Echo, pesan sukses, data yang akan dikirim, dan kode status HTTP, lalu mengembalikan response JSON dengan format yang sudah ditentukan
func NewSuccessResponse(c echo.Context, message string, data interface{}, statusCode int) error {
	return c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"data":    data,
		"status":  statusCode,
	})
}
