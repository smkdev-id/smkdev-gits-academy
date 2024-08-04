package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/FuadGrimaldi/bookstore-api/pkg/config"
	"github.com/FuadGrimaldi/bookstore-api/pkg/routes"
	"github.com/labstack/echo/v4"
)

// Structur Server
type Server struct {
	*echo.Echo
	cfg *config.Config
}

// Initializes a new Echo server instance
func NewServer(cfg *config.Config, bookRoutes []*routes.Route) *Server {
	e := echo.New()
	for _, v := range bookRoutes {
		e.Add(v.Method, v.Path, v.Handler)
	}
	return &Server{e, cfg}
}

// Starts the Echo server
func (s *Server) Run() {
	go func ()  {
		err := s.Start(fmt.Sprintf(":%s", s.cfg.Port))
		log.Fatal(err)
	} ()
}

// Handles graceful shutdown of the server
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}