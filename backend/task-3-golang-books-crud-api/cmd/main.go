package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.InitDB()
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
