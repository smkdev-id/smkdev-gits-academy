package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupMiddleware mengatur middleware untuk aplikasi Echo.
func SetupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())  // Middleware untuk mencatat log setiap permintaan
	e.Use(middleware.Recover()) // Middleware untuk menangani panic dan pemulihan dari kesalahan
	e.Use(middleware.CORS())    // Middleware untuk menangani Cross-Origin Resource Sharing (CORS)
}
