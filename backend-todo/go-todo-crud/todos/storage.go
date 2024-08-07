package todos

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

	var db *sql.DB

	func InitDB(filepath string)  {
		var err error
		db, err = sql.Open("sqlite3", filepath)
		if err != nil{
			log.Fatal(err)
		}
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
		createTable()
	}

	func createTable() {
		createTableSql := `CREATE TABLE IF NOT EXISTS todos (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"status" TEXT
		);`

		statement, err := db.Prepare(createTableSql)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
	}