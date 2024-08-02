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
		FindAll(ctx context.Context, page, pageSize int) ([]*response.BookResponse, int, error)
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

	if err := s.bookRepository.Create(ctx, newBook); err != nil {
		return err
	}

	return nil
}

// Delete implements BookService.
func (s *bookService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements BookService.
func (s *bookService) FindAll(ctx context.Context, page, pageSize int) ([]*response.BookResponse, int, error) {
	books, totalRecords, err := s.bookRepository.FindAll(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var bookResponses []*response.BookResponse

	for _, book := range books {
		updatedAtStr := ""
		if book.UpdatedAt != nil {
			updatedAtStr = book.UpdatedAt.String()
		}

		bookResponses = append(bookResponses, &response.BookResponse{
			Id:                book.ID,
			Isbn:              book.ISBN,
			Title:             book.Title,
			Description:       book.Description,
			Author:            book.Author,
			YearOfManufacture: book.YearOfManufacture,
			Stock:             book.Stock,
			Price:             book.Price,
			CreatedAt:         book.CreatedAt.String(),
			UpdatedAt:         updatedAtStr,
		})
	}

	return bookResponses, int(totalRecords), nil
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
