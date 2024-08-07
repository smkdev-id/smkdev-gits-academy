package server

import (
	"EkoEdyPurwanto/task-3/config"
	"EkoEdyPurwanto/task-3/internal/adapter/manager"
	transport "EkoEdyPurwanto/task-3/internal/transport/http"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		UseCaseM manager.UseCaseManager
		Engine   *fiber.App
		Host     string
		Log      *logrus.Logger
	}
)

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.Engine.Listen(s.Host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	// init middleware here
}

func (s *Server) initControllers() {
	transport.NewBookTransport(s.UseCaseM.BookUseCase(), s.Engine).Route()
}

// NewServer Constructor
func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	var (
		// instance repository
		repositoryManager = manager.NewRepositoryManager(infraManager)

		// instance useCase
		validate       = validator.New()
		logger         = logrus.New()
		useCaseManager = manager.NewUseCaseManager(repositoryManager, validate, logger)

		hostAndPort = fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)

		app = fiber.New()
	)
	return &Server{
		UseCaseM: useCaseManager,
		Engine:   app,
		Host:     hostAndPort,
		Log:      logger,
	}
}
