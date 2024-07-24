package handlers

import (
	"database/sql"
	"net/http"
	"go-crud-api/models"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var task models.Task

		c.Bind(&task)

		id, err := models.PutTask(db, task.Name, task.Status)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func EditTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var task models.Task
		c.Bind(&task)

		_, err := models.EditTask(db, task.ID, task.Name, task.Status)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": task,
			})
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteTask(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}

	}
} 