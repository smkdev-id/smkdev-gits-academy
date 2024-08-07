package main

import (
    "smkdev-task-3/pkg/config"
    "smkdev-task-3/pkg/routes"
)

func main() {
    config.InitDB()
    e := routes.InitRoutes()
    e.Logger.Fatal(e.Start(":8080"))
}
