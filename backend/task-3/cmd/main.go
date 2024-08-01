package main

import (
	"context"
	"log"

	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/builders"
	"github.com/alwiirfan/pkg/database"
	"github.com/alwiirfan/pkg/server"
)

func main() {

	ctx := context.Background()

	// load configs
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configs: %v", err)
	}

	// initialize sqlite database connection
	db, err := database.SqliteConnection(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to sqlite database: %v", err)
	}

	// build public routes
	publicRoutes := builders.BuildPublicRoutes(cfg, db)

	// start server
	server := server.NewServer(cfg, publicRoutes, nil)
	// run server
	server.Run()
	// graceful shutdown server
	server.GracefulShutdown()
}
