package main

import (
	"golang-crud/db"
	"golang-crud/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "todo.db"
)


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := db.GetConnection(DB_NAME)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	handlers.Route(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}