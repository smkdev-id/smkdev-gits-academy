package delivery

import (
	"EkoEdyPurwanto/task-2/config"
	"EkoEdyPurwanto/task-2/delivery/controller"
	"EkoEdyPurwanto/task-2/manager"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		UseCaseM manager.UseCaseManager
		Engine   *http.ServeMux
		Host     string
		Log      *logrus.Logger
	}
)

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	fmt.Println("Starting server at", s.Host)
	err := http.ListenAndServe(s.Host, s.Engine)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	// init middleware here
}

func (s *Server) initControllers() {
	controller.NewTodoListController(s.UseCaseM.TodoListUseCase(), s.Engine).Route()
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

		app = http.NewServeMux()
	)
	return &Server{
		UseCaseM: useCaseManager,
		Engine:   app,
		Host:     hostAndPort,
		Log:      logger,
	}
}
