package main

import (
	"BookStore/pkg/config"
	"BookStore/pkg/routes"
	"log"
)

func main() {

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := routes.SetupRouter(db)

	routes.RunServer(router)
}
