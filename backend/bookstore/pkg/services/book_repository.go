package services

import (
	"bookstore/pkg/helper"
	"bookstore/pkg/models"
	"bookstore/pkg/repository"
	"time"

	"github.com/gofrs/uuid"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(request helper.Request) (models.Book, error)
	UpdateBook(id string, request helper.Request) (models.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	books, err := s.repository.GetAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *bookService) GetBookByID(id string) (models.Book, error) {
	book, err := s.repository.GetByID(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *bookService) CreateBook(request helper.Request) (models.Book, error) {
	var book models.Book
	id, _ := uuid.NewV4()
	book.ID = id.String()
	book.Title = request.Title
	book.Description = request.Description
	book.Genre = request.Genre
	book.ISBN = request.ISBN
	book.Author = request.Author
	book.Price = request.Price
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	book, err := s.repository.Create(book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *bookService) UpdateBook(id string, request helper.Request) (models.Book, error) {
	book, err := s.repository.GetByID(id)
	if err != nil {
		return book, err
	}

	book.Title = request.Title
	book.Description = request.Description
	book.Genre = request.Genre
	book.ISBN = request.ISBN
	book.Author = request.Author
	book.Price = request.Price
	book.UpdatedAt = time.Now()

	newbook, err := s.repository.Update(book)
	if err != nil {
		return newbook, err
	}

	return newbook, nil
}

func (s *bookService) DeleteBook(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
