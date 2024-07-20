package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Todo     string `json:"todo" gorm:"type:text; not null"`
	Finished bool   `json:"finished" gorm:"type:not null"`
}

func main() {
	// Initialize ECHO
	e := echo.New()
	// Initialize Database
	sqliteDatabase, err := gorm.Open(sqlite.Open("database-todo.db"), &gorm.Config{})
	if err != nil {
		log.Println("Database Not Found. Creating one....")
		file, err := os.Create("database-todo.db") // Create SQLite file
		if err != nil {
			log.Println("Failed to create a database")
		}
		file.Close()
	}
	// Table Migrate
	sqliteDatabase.AutoMigrate(&Todo{})

	// Routes
	e.GET("/", fetchAllTodos(sqliteDatabase))
	e.POST("/", createTodo(sqliteDatabase))
	e.PUT("/:id", updateTodo(sqliteDatabase))
	e.DELETE("/:id", deleteTodo(sqliteDatabase))

	// Logger Server
	e.Logger.Fatal(e.Start(":8000"))
}

// Routes HandlerFunc
func fetchAllTodos(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todoData []Todo
		result := db.Find(&todoData)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to fetch toods",
				"error":   result.Error.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    todoData,
		})
	}
}

func createTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request Todo
		if err := c.Bind(&request); err != nil {
			return err
		}

		newTodo := new(Todo)
		newTodo.Todo = request.Todo
		newTodo.Finished = false

		result := db.Create(&newTodo)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to insert data",
				"error":   result.Error.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    newTodo,
		})

	}
}

func updateTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")
		var request Todo

		intVar, errid := strconv.Atoi(id)

		if err := c.Bind(&request); err != nil {
			return err
		}

		if errid != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "invalid id.",
			})
		}

		var todoData []Todo

		db.First(&todoData, intVar)

		if len(todoData) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "id not found.",
			})
		}

		db.Save(Todo{ID: int64(intVar), Todo: todoData[len(todoData)-1].Todo, Finished: request.Finished})

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})

	}
}

func deleteTodo(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")
		intVar, err := strconv.Atoi(id)
		var todos []Todo

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "invalid id.",
			})
		}

		db.First(&todos, intVar)

		if len(todos) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "id not found.",
			})
		}

		db.Delete(&Todo{}, intVar)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})
	}
}
