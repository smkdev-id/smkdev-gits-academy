package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alwi09/config"
	"github.com/alwi09/controller"
	"github.com/alwi09/manager"
	"github.com/alwi09/repository"
	"github.com/alwi09/router"
	"github.com/alwi09/service"
)

func StartServer() {

	// itinialize configuration for database
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// itinialize database
	mysqlManager, err := manager.NewMysqlManager(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// initialize repository, service and controller
	repository := repository.NewTodoRepository(mysqlManager.Conn())
	service := service.NewTodoService(repository)
	controller := controller.NewTodoController(service)

	// intialize router
	router := router.NewRouter(controller)

	fmt.Printf("Server is running on port %s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(os.Getenv("APP_PORT"), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
