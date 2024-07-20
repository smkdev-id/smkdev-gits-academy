package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection(dbname string)(*sql.DB, error) {
	// init SQLite3 database
	db, err := sql.Open("sqlite3", dbname) 
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %s", err)
	}
	fmt.Println("Connected successfully to the database")
	return db, nil
}