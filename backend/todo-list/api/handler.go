package api

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
	"github.com/PorcoGalliard/Todo-List/db"
    "github.com/PorcoGalliard/Todo-List/types"
)

func GetTodos(c echo.Context) error {
    rows, err := db.DB.Query("SELECT id, title, completed FROM todos")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    defer rows.Close()

    todos := []types.Todo{}
    for rows.Next() {
        var todo types.Todo
        err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, err.Error())
        }
        todos = append(todos, todo)
    }

    return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
    todo := new(types.Todo)
    if err := c.Bind(todo); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    result, err := db.DB.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    id, _ := result.LastInsertId()
    todo.ID = int(id)
    return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo := new(types.Todo)
    if err := c.Bind(todo); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    _, err := db.DB.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusNoContent)
}

func DeleteTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    _, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusNoContent)
}
