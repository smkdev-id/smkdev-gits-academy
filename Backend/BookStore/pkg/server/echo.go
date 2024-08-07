package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server adalah struct yang meng-embed Echo
type Server struct {
	*echo.Echo
}

// NewServer membuat instance baru dari Server
func NewServer() *Server {
	e := echo.New() // Buat instance Echo baru

	// Pasang middleware
	e.Use(
		middleware.Logger(),  // Middleware untuk logging
		middleware.Recover(), // Middleware untuk recovery dari panic
		middleware.CORS(),    // Middleware untuk mengizinkan CORS
	)

	// Kembalikan instance Server dengan Echo yang telah dikonfigurasi
	return &Server{e}
}
