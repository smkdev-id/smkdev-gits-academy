package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alwi09/config"
	"github.com/alwi09/manager"
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

	fmt.Println(mysqlManager)

	port := ":8080"
	fmt.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
