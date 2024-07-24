package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
    var err error
    DB, err = sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    createTable()
}

func createTable() {
    createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
        "title" TEXT,
        "completed" BOOLEAN
    );`

    statement, err := DB.Prepare(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    statement.Exec()
}
