package handlers

import (
	"database/sql"
	"fmt"
	"golang-crud/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
desc 	: handler to get all the data todos
method 	: GET
*/
func getAllTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := service.GetAllTodos(db)
		fmt.Println(todos)
		if err != nil {
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "successfully get all todos", todos)
	}
}

/*
desc 	: handler to get todo by id
method 	: GET 
*/
func getTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		t, err := service.GetTodoById(db, id)
		if err == sql.ErrNoRows {
			return JSONResponse(c, http.StatusNotFound, err.Error(), nil)
		} else if err != nil {
			c.Logger().Error(err) 
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "successfully get todo by id", t)
	}
}


/*
desc 	: handler to create todo
method 	: POST
*/
func createTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(service.Todo)
		if err := c.Bind(t); err != nil {
			c.Logger().Error(err)
			return JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		}
		if err :=t.Create(db); err != nil {
			c.Logger().Error(err)
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusCreated, "Create successfully",t)
	}
}

/*
desc 	: handler to update todo
method 	: PUT
*/
func updateTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		t := new(service.Todo)
		if err := c.Bind(t); err != nil {
			c.Logger().Error()
			return JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		} 
		t.Id = id
		if err := t.Update(db); err != nil {
			c.Logger().Error()
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "Update successfully",t)
	}
}

/*
desc 	: handler to update todo
method 	: DELETE
*/
func deleteTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := service.DeleteTodoById(db, id); err != nil {
			c.Logger().Error()
			return JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return JSONResponse(c, http.StatusOK, "Delete successfully",nil)
	}
}

