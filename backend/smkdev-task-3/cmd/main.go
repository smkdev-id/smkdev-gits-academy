package main

import (
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    config.InitDB()
    routes.InitRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
