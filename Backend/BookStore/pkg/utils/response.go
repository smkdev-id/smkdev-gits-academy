package utils

import (
	"github.com/labstack/echo/v4"
)

// Response adalah struktur data untuk format respons API.
type Response struct {
	Status  string      `json:"status"`      // Status respons (misalnya, "success" atau "error")
	Message string      `json:"message"`     // Pesan yang menyertai respons
	Data    interface{} `json:"data,omitempty"` // Data respons (opsional, hanya ditampilkan jika ada)
}

// SuccessResponse mengirimkan respons sukses dengan data.
func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, Response{
		Status:  "success",
		Message: "Request was successful",
		Data:    data,
	})
}

// SuccessResponseWithoutData mengirimkan respons sukses tanpa data.
func SuccessResponseWithoutData(c echo.Context, code int) error {
	return c.JSON(code, Response{
		Status:  "success",
		Message: "Request was successful",
	})
}

// ErrorResponse mengirimkan respons kesalahan dengan pesan.
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Status:  "error",
		Message: message,
	})
}

// ErrorResponseValidation mengirimkan respons kesalahan dengan pesan validasi dan data kesalahan.
func ErrorResponseValidation(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, Response{
		Status:  "error",
		Message: "Validation error",
		Data:    data,
	})
}
