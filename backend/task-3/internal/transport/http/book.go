package transport

import (
	"EkoEdyPurwanto/task-3/internal/model/dto/req"
	"EkoEdyPurwanto/task-3/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type (
	BookTransport struct {
		UseCase usecase.BookUseCase
		Engine  *fiber.App
	}
)

func NewBookTransport(useCase usecase.BookUseCase, engine *fiber.App) *BookTransport {
	return &BookTransport{
		UseCase: useCase,
		Engine:  engine,
	}
}

// Route sets up the routes for the Book API
func (b *BookTransport) Route() {
	api := b.Engine.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/book", b.createHandler)
	v1.Get("/book", b.getAllHandler)
	v1.Post("/book/getById", b.getByIdHandler)
	v1.Put("/book", b.updateHandler)
	v1.Delete("/book", b.deleteHandler)
}

// createHandler handles book creation requests
func (b *BookTransport) createHandler(ctx *fiber.Ctx) error {
	var request req.CreateRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	if err := b.UseCase.Create(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Book created successfully"})
}

// getAllHandler handles requests to fetch all books
func (b *BookTransport) getAllHandler(ctx *fiber.Ctx) error {
	books, err := b.UseCase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(books)
}

// getByIdHandler handles requests to fetch a book by ID
func (b *BookTransport) getByIdHandler(ctx *fiber.Ctx) error {
	var request struct {
		Id string `json:"id" validate:"required"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	book, err := b.UseCase.GetById(request.Id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "Book not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(book)
}

// updateHandler handles requests to update a book
func (b *BookTransport) updateHandler(ctx *fiber.Ctx) error {
	var request req.UpdateRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	if err := b.UseCase.UpdateStatus(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book updated successfully"})
}

// deleteHandler handles requests to delete a book
func (b *BookTransport) deleteHandler(ctx *fiber.Ctx) error {
	var request struct {
		Id string `json:"id" validate:"required"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Error parsing request body",
		})
	}

	err := b.UseCase.Delete(req.DeleteRequest{Id: request.Id})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}
