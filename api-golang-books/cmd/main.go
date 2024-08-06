package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/config"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/routes"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/utils"
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
