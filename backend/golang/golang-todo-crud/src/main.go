package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID         uuid.UUID `json:"id"`
	Todo       string    `json:"todo"`
	IsDone     bool      `json:"isDone"`
	CreatedAt  string    `json:"createdAt"`
	UpdatedAt  string    `json:"updatedAt"`
	FinishedAt string    `json:"finishedAt"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		panic(err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id TEXT PRIMARY KEY,
        todo TEXT,
        is_done BOOLEAN DEFAULT FALSE,
        created_at TEXT,
        updated_at TEXT,
        finished_at TEXT DEFAULT ''
    );`

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	initDB()

	e.GET("/todos", getTodos)
	e.GET("/todos/:id", getTodo)
	e.POST("/todos", createTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}

func getTodos(c echo.Context) error {
	rows, err := db.Query("SELECT id, todo, is_done, created_at, updated_at, finished_at FROM todos")
	if err != nil {
		return err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Todo, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt, &todo.FinishedAt); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    todos,
	})
}

func getTodo(c echo.Context) error {
	id := c.Param("id")
	var todo Todo
	if err := db.QueryRow("SELECT id, todo, is_done, created_at, updated_at, finished_at FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Todo, &todo.IsDone, &todo.CreatedAt, &todo.UpdatedAt, &todo.FinishedAt); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Todo not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    todo,
	})
}

func createTodo(c echo.Context) error {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return err
	}

	if todo.Todo == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Field 'todo' is required",
		})
	}

	todo.ID = uuid.New()
	todo.CreatedAt = time.Now().Format(time.RFC3339)
	todo.UpdatedAt = todo.CreatedAt

	todo.IsDone = false
	todo.FinishedAt = ""

	_, err := db.Exec("INSERT INTO todos (id, todo, is_done, created_at, updated_at, finished_at) VALUES (?, ?, ?, ?, ?, ?)",
		todo.ID, todo.Todo, todo.IsDone, todo.CreatedAt, todo.UpdatedAt, todo.FinishedAt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "successfully created a Todo",
		"data":    todo,
	})
}

func updateTodo(c echo.Context) error {
	id := c.Param("id")

	var existingTodo Todo
	err := db.QueryRow("SELECT id, todo, is_done, created_at, updated_at, finished_at FROM todos WHERE id = ?", id).Scan(
		&existingTodo.ID, &existingTodo.Todo, &existingTodo.IsDone, &existingTodo.CreatedAt, &existingTodo.UpdatedAt, &existingTodo.FinishedAt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Todo not found",
		})
	}

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return err
	}

	if todo.Todo == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Field 'todo' is required",
		})
	}
	if todo.IsDone == existingTodo.IsDone {
		todo.IsDone = existingTodo.IsDone
	} else {
		if todo.IsDone {
			todo.FinishedAt = time.Now().Format(time.RFC3339)
		} else {
			todo.FinishedAt = ""
		}
	}

	todo.ID = existingTodo.ID
	todo.CreatedAt = existingTodo.CreatedAt
	todo.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err = db.Exec("UPDATE todos SET todo = ?, is_done = ?, updated_at = ?, finished_at = ? WHERE id = ?",
		todo.Todo, todo.IsDone, todo.UpdatedAt, todo.FinishedAt, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully updated a Todo",
		"data":    todo,
	})
}

func deleteTodo(c echo.Context) error {
	id := c.Param("id")

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM todos WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "An error occurred",
		})
	}

	if !exists {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Todo not found",
		})
	}

	_, err = db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "An error occurred",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted the Todo",
	})
}
