package app

import (
	"database/sql"

	"golang-bookstore/helper"

	_ "github.com/mattn/go-sqlite3"
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
