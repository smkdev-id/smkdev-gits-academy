package database

import (
	"database/sql"
    _ "modernc.org/sqlite"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "database/todos.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	return db, nil
}

func CreateTodoTable(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            description TEXT,
            completed BOOLEAN,
            created_at TIMESTAMP,
            updated_at TIMESTAMP
        )
    `)
    if err != nil {
        return fmt.Errorf("failed to create table: %v", err)
    }
    return nil
}
