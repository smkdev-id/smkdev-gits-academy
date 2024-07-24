// dbutils.go
package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
    db   *sql.DB
    lock = sync.Mutex{}
)

func initDB(dataSourceName string) {
    var err error
    db, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    // Create todo table if it doesn't exist
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS todo (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL
    );`)
    if err != nil {
        log.Fatal(err)
    }
}

func closeDB() {
    db.Close()
}
