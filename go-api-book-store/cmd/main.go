package main

import (
	"fmt"

	"github.com/hasban-fardani/go-api-book-store/pkg/config"
	"github.com/hasban-fardani/go-api-book-store/pkg/routes"
	"github.com/hasban-fardani/go-api-book-store/pkg/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	// load config from .env file
	cfg, err := config.NewConfig(".env")
	utils.HandleError(err)

	server := echo.New()
	routes.BookRoute(server)

	err = server.Start(":" + cfg.Port)
	fmt.Println(err.Error())
}
