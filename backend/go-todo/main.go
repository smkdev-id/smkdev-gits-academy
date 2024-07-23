package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Todo struct represents a task with a title and status.
type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"` // Unique identifier for the todo.
	Title  string `json:"title"`                // Title of the todo.
	Status string `json:"status"`               // Status of the todo (e.g., pending, completed).
}

var DB *gorm.DB

// initDatabase initializes the database and performs migration.
func initDatabase() {
	var err error
	// Open the SQLite database using the GORM driver.
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate the schema, ensuring the table is created.
	DB.AutoMigrate(&Todo{})
}

// fetchTodos handles GET requests to fetch all todos.
func fetchTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos) // Retrieve all todos from the database.
	c.JSON(http.StatusOK, todos)
}

// createTodo handles POST requests to create a new todo.
func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		// If there's an error in binding JSON, return a bad request response.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create the new todo in the database.
	DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

// updateTodo handles PUT requests to update an existing todo.
func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	// Find the todo by its ID.
	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated todo to the database.
	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// deleteTodo handles DELETE requests to remove a todo.
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	// Delete the todo from the database by its ID.
	if err := DB.Delete(&Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Todo deleted"})
}

func main() {
	initDatabase() // Initialize the database.

	r := gin.Default()

	// Define the routes and handlers.
	r.GET("/todos", fetchTodos)        // Fetch all todos.
	r.POST("/todos", createTodo)       // Create a new todo.
	r.PUT("/todos/:id", updateTodo)    // Update an existing todo.
	r.DELETE("/todos/:id", deleteTodo) // Delete a todo.

	// Run the server on port 8080.
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
