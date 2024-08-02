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
		FindAllSearch(ctx context.Context, page, pageSize int, req *request.SearchBookRequest) ([]*response.BookResponse, int, error)
		FindByID(ctx context.Context, id string) (*response.BookResponse, error)
		FindByISBN(ctx context.Context, isbn string) (*response.BookResponse, error)
		Create(ctx context.Context, req *request.CreateBookRequest) error
		Update(ctx context.Context, id string, req *request.UpdateBookRequest) error
		DeleteByID(ctx context.Context, id string) error
		DeleteAll(ctx context.Context) error
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
func (s *bookService) DeleteByID(ctx context.Context, id string) error {
	existingBook, err := s.bookRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.bookRepository.DeleteByID(ctx, existingBook.ID)
}

// DeleteAll implements BookService.
func (s *bookService) DeleteAll(ctx context.Context) error {
	if err := s.bookRepository.DeleteAll(ctx); err != nil {
		return err
	}

	return nil
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

// FindAllSearch implements BookService.
func (s *bookService) FindAllSearch(ctx context.Context, page, pageSize int, req *request.SearchBookRequest) ([]*response.BookResponse, int, error) {
	books, totalRecords, err := s.bookRepository.FindAllSearch(ctx, page, pageSize, req)
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
	existingBook, err := s.bookRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedAtStr := ""
	if existingBook.UpdatedAt != nil {
		updatedAtStr = existingBook.UpdatedAt.String()
	}

	bookResp := &response.BookResponse{
		Id:                existingBook.ID,
		Isbn:              existingBook.ISBN,
		Title:             existingBook.Title,
		Description:       existingBook.Description,
		Author:            existingBook.Author,
		YearOfManufacture: existingBook.YearOfManufacture,
		Stock:             existingBook.Stock,
		Price:             existingBook.Price,
		CreatedAt:         existingBook.CreatedAt.String(),
		UpdatedAt:         updatedAtStr,
	}

	return bookResp, nil
}

// FindByISBN implements BookService.
func (s *bookService) FindByISBN(ctx context.Context, isbn string) (*response.BookResponse, error) {
	existingBook, err := s.bookRepository.FindByISBN(ctx, isbn)
	if err != nil {
		return nil, err
	}

	updatedAtStr := ""
	if existingBook.UpdatedAt != nil {
		updatedAtStr = existingBook.UpdatedAt.String()
	}

	bookResp := &response.BookResponse{
		Id:                existingBook.ID,
		Isbn:              existingBook.ISBN,
		Title:             existingBook.Title,
		Description:       existingBook.Description,
		Author:            existingBook.Author,
		YearOfManufacture: existingBook.YearOfManufacture,
		Stock:             existingBook.Stock,
		Price:             existingBook.Price,
		CreatedAt:         existingBook.CreatedAt.String(),
		UpdatedAt:         updatedAtStr,
	}

	return bookResp, nil
}

// Update implements BookService.
func (s *bookService) Update(ctx context.Context, id string, req *request.UpdateBookRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return err
	}

	existingBook, err := s.bookRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	//init updated at because type *time.Time is nil
	updatedAt := time.Now().Local()

	// update book
	existingBook.Title = req.Title
	existingBook.Description = req.Description
	existingBook.Author = req.Author
	existingBook.YearOfManufacture = req.YearOfManufacture
	existingBook.Stock = req.Stock
	existingBook.Price = req.Price
	existingBook.UpdatedAt = &updatedAt

	return s.bookRepository.Update(ctx, existingBook)
}

func NewBookService(cfg *configs.Config, bookRepository repositories.BookRepository) BookService {
	return &bookService{
		cfg:            cfg,
		bookRepository: bookRepository,
		validate:       validator.New(),
	}
}
