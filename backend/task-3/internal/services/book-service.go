package services

import (
	"context"
	"time"

	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/dto/response"
	"github.com/alwiirfan/internal/entities"
	"github.com/alwiirfan/internal/repositories"
	"github.com/alwiirfan/internal/utils"
	"github.com/go-playground/validator/v10"
)

type (
	BookService interface {
		FindAll(ctx context.Context) ([]*response.BookResponse, error)
		FindByID(ctx context.Context, id string) (*response.BookResponse, error)
		Create(ctx context.Context, req *request.CreateBookRequest) error
		Update(ctx context.Context, req *request.UpdateBookRequest) error
		Delete(ctx context.Context, id string) error
	}

	bookService struct {
		cfg            *configs.Config
		bookRepository repositories.BookRepository
		validate       *validator.Validate
	}
)

// Create implements BookService.
func (s *bookService) Create(ctx context.Context, req *request.CreateBookRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return err
	}

	newBook := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             req.Title,
		Description:       req.Description,
		Author:            req.Author,
		YearOfManufacture: req.YearOfManufacture,
		Stock:             req.Stock,
		Price:             req.Price,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	return s.bookRepository.Create(ctx, newBook)
}

// Delete implements BookService.
func (s *bookService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements BookService.
func (s *bookService) FindAll(ctx context.Context) ([]*response.BookResponse, error) {
	panic("unimplemented")
}

// FindByID implements BookService.
func (s *bookService) FindByID(ctx context.Context, id string) (*response.BookResponse, error) {
	panic("unimplemented")
}

// Update implements BookService.
func (s *bookService) Update(ctx context.Context, req *request.UpdateBookRequest) error {
	panic("unimplemented")
}

func NewBookService(cfg *configs.Config, bookRepository repositories.BookRepository) BookService {
	return &bookService{
		cfg:            cfg,
		bookRepository: bookRepository,
		validate:       validator.New(),
	}
}
