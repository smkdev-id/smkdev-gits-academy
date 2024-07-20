package main

import (
	"golang-crud/db"
	"golang-crud/handlers"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// database name
const DB_NAME    string = "todo.db"


func main() {
	// echo initialization
	e := echo.New()

	// server connection to sqlite3 database
	db, err := db.GetConnection(DB_NAME)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// route initialization
	handlers.Route(e, db)

	// Echo server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}