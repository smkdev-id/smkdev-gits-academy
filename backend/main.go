package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type todo struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
}

func createTodo(c echo.Context) error {
    lock.Lock()
    defer lock.Unlock()

    t := new(todo)
    if err := c.Bind(t); err != nil {
        return err
    }

    stmt, err := db.Prepare("INSERT INTO todo (task) VALUES (?)")
    if err != nil {
        return err
    }
    res, err := stmt.Exec(t.Task)
    if err != nil {
        return err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    t.ID = int(id)

    return c.JSON(http.StatusCreated, t)
}

func getTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    var t todo
    err := db.QueryRow("SELECT id, task FROM todo WHERE id = ?", id).Scan(&t.ID, &t.Task)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, t)
}

func updateTodo(c echo.Context) error {
    lock.Lock()
    defer lock.Unlock()

    t := new(todo)
    if err := c.Bind(t); err != nil {
        return err
    }

    id, _ := strconv.Atoi(c.Param("id"))

    stmt, err := db.Prepare("UPDATE todo SET task = ? WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(t.Task, id)
    if err != nil {
        return err
    }

    t.ID = id
    return c.JSON(http.StatusOK, t)
}

func deleteTodo(c echo.Context) error {
    lock.Lock()
    defer lock.Unlock()

    id, _ := strconv.Atoi(c.Param("id"))

    stmt, err := db.Prepare("DELETE FROM todo WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    return c.NoContent(http.StatusNoContent)
}

func getAllTodos(c echo.Context) error {
    rows, err := db.Query("SELECT id, task FROM todo")
    if err != nil {
        return err
    }
    defer rows.Close()

    todos := []todo{}
    for rows.Next() {
        var t todo
        if err := rows.Scan(&t.ID, &t.Task); err != nil {
            return err
        }
        todos = append(todos, t)
    }

    return c.JSON(http.StatusOK, todos)
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbFile := os.Getenv("DB_FILE")
    initDB(dbFile)
    defer closeDB()

    port := os.Getenv("PORT")

    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/todos", getAllTodos)
    e.POST("/todos", createTodo)
    e.GET("/todos/:id", getTodo)
    e.PUT("/todos/:id", updateTodo)
    e.DELETE("/todos/:id", deleteTodo)

    // Start server
    e.Logger.Fatal(e.Start(":" + port))
}
