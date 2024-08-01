package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/http/routes"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	cfg *configs.Config
}

func NewServer(cfg *configs.Config, publicRoutes, privateRoutes []*routes.Route) *Server {
	e := echo.New()
	for _, v := range publicRoutes {
		e.Add(v.Method, v.Path, v.Handler)
	}

	for _, v := range privateRoutes {
		e.Add(v.Method, v.Path, v.Handler) // with MiddleWareFunc
	}

	return &Server{e, cfg}
}

func (s *Server) Run() {
	go func() {
		err := s.Start(fmt.Sprintf(":%s", s.cfg.AppPort))
		log.Fatal(err)
	}()
}

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
