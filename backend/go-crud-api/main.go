package main

import (
	"database/sql"
	"go-crud-api/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)

	// daftar api
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PutTask(db))
	e.PUT("/tasks", handlers.EditTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}