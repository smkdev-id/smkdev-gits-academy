package main

import (
	db "go_crud_todos/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", index)

	app.Post("/", create)

	app.Put("/:id", update)

	app.Delete("/:id", delete)

	app.Listen(":3000")
}

func index(c *fiber.Ctx) error {
	todos := db.GetTodos()
	return c.JSON(todos)
}

func create(c *fiber.Ctx) error {
	var todo *db.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.CreateTodo(todo)

	return c.JSON(todo)
}

func update(c *fiber.Ctx) error {
	var id = c.Params("id")

	var todo *db.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.UpdateTodo(id, todo)
	return c.JSON(todo)
}

func delete(c *fiber.Ctx) error {
	var id = c.Params("id")

	err := db.DeleteTodo(id)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}
