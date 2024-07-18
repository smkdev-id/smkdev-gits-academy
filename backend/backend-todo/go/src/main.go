package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID       int64  `json:"id"`
	Todo     string `json:"todo"`
	Finished bool   `json:"finished"`
}

func main() {
	// Initialize ECHO
	e := echo.New()
	// Initialize Database
	sqliteDatabase, err := sql.Open("sqlite3", "./database-todo.db")
	if err != nil {
		log.Println("Database Not Found. Creating one....")
		file, err := os.Create("database-todo.db") // Create SQLite file
		if err != nil {
			log.Println("Failed to create a database")
		}
		file.Close()
	}
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)

	// Insert Record(debug)
	insertTodo(sqliteDatabase, "Menggoreng Nasi", false)
	insertTodo(sqliteDatabase, "Menggoreng Kayu", true)

	// Routes
	e.GET("/", fetchAllTodos(sqliteDatabase))

	// Logger Server
	e.Logger.Fatal(e.Start(":8000"))
}

// Func Create Table If Not Exist
func createTable(db *sql.DB) {
	createTodoQuery := `CREATE TABLE IF NOT EXISTS todo (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"todo" TEXT,
		"finished" BOOL	
	  );`

	log.Println("Create todo table if not exist...")
	statement, err := db.Prepare(createTodoQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Todo Table Is Successfully Created")
}

// Func for fetch all todos
func fetchAllTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var obj Todo
		var arrobj []Todo

		row, err := db.Query("SELECT * FROM todo")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer row.Close()
		for row.Next() {

			err = row.Scan(&obj.ID, &obj.Todo, &obj.Finished)

			if err != nil {
				log.Println(err)
			}

			arrobj = append(arrobj, obj)

		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success",
			"data":    arrobj,
		})
	}

}

func insertTodo(db *sql.DB, todo string, finished bool) {
	log.Println("Inserting todo record ...")
	insertTodo := `INSERT INTO todo(todo, finished) VALUES (?, ?)`
	statement, err := db.Prepare(insertTodo)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(todo, finished)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
