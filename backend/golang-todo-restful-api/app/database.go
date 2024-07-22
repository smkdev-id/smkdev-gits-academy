package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database/todo.db")
	helper.PanicIfError(err)

	// Buat tabel
	sqlStmt := `
		 CREATE TABLE IF NOT EXISTS todolist (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			description TEXT,
			completed BOOLEAN
    );
    `

	_, err = db.Exec(sqlStmt)
	helper.PanicIfError(err)

	return db
}
