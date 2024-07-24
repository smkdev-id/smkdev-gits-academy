package todos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	rows, err := db.Query("SELECT id, title, status FROM todos")
	if err != nil {
		return err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next(){
		var todo Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return err
	}

	statement, err := db.Prepare("INSERT INTO todos (title, status) VALUES (?, ?)")
	if err != nil{
		return err
	}

	result, err := statement.Exec(todo.Title, todo.Status)
	if err != nil{
		return err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return err
	}

	todo.ID = int(id)

	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	var todo Todo
	id,_ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&todo); err != nil{
		return err
	}

	statement, err := db.Prepare("UPDATE todos SET title = ?, status = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(todo.Title, todo.Status, id)
	if err != nil {
		return err
	}

	todo.ID = id
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	statement, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}