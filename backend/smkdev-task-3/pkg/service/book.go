package service

import (
	"BookStore/pkg/models"
	"BookStore/pkg/repository"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookService interface {
	GetBooks(c *gin.Context) ([]models.Book, error)
	GetBooksByID(c *gin.Context, id uuid.UUID) (*models.Book, error)
	CreateBook(c *gin.Context, request *models.CreateBookRequest) (*models.Book, error)
	UpdateBook(c *gin.Context, request *models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(c *gin.Context, id uuid.UUID) error
}

type bookService struct {
	// db         *gorm.DB
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{repository: repository}
}

func (b *bookService) GetBooks(c *gin.Context) ([]models.Book, error) {
	books, err := b.repository.FindAll(c)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (b *bookService) GetBooksByID(c *gin.Context, id uuid.UUID) (*models.Book, error) {
	// idbook, err := uuid.Parse(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
	// 	return nil, err
	// }
	books, err := b.repository.FindByID(c, id)
	if err != nil {
		return nil, err
	}
	// var books models.Book

	// if err := b.db.First(&books).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Printf("Books error: %+v", err)
	// 	return nil, err
	// }

	// if err := b.db.First(&books).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	log.Printf("Books error: %+v", err)
	// 	return nil, err
	// }

	// Log the retrieved books and return them
	log.Printf("Books retrieved: %+v", books)
	return books, nil

}

func (b *bookService) CreateBook(c *gin.Context, request *models.CreateBookRequest) (*models.Book, error) {
	book := models.Book{
		ID:        uuid.New(),
		Title:     request.Title,
		Publisher: request.Publisher,
		Author:    request.Author,
		Price:     request.Price,
		Stock:     request.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := b.repository.CreateBook(c, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookService) UpdateBook(c *gin.Context, request *models.UpdateBookRequest) (*models.Book, error) {
	id, err := uuid.Parse(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return nil, err
	}
	book, err := b.repository.FindByID(c, id)
	if err != nil {
		return nil, err
	}
	if request.Title != "" {
		book.Title = request.Title
	}

	if request.Author != "" {
		book.Author = request.Author
	}
	if request.Publisher != "" {
		book.Publisher = request.Publisher
	}
	if strconv.Itoa(request.Stock) != "" {
		book.Stock = request.Stock
	}
	if strconv.Itoa(request.Price) != "" {
		book.Price = request.Price
	}
	book.UpdatedAt = time.Now()
	// book := models.Book{
	// 	// ID:        request.ID,
	// 	Title:     request.Title,
	// 	Publisher: request.Publisher,
	// 	Author:    request.Author,
	// 	Price:     request.Price,
	// 	UpdatedAt: time.Now(),
	// }
	err = b.repository.UpdateBook(c, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (b *bookService) DeleteBook(c *gin.Context, id uuid.UUID) error {
	return b.repository.DeleteBook(c, id)
}
