package main

import (
	"todolist/config"
	"todolist/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	//inisialisasi koneksi ke database
	config.Connect()

	//membuat instance baru dari echo
	e := echo.New()

	//inisialisasi route todolist 
	routes.Init(e)

	//server web berjalan di port 3030
	e.Logger.Fatal(e.Start(":3030"))

}
