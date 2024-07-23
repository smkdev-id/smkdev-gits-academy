package main

import (
	"todos/database"
	"todos/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	err = database.CreateTodoTable(db)
	if err != nil {
		panic("failed to create table")
	}

	router.TodoRouter(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
