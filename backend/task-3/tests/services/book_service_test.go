package services

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/entities"
	"github.com/alwiirfan/internal/services"
	"github.com/alwiirfan/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BookServiceSuite struct {
	suite.Suite
	cfg         *configs.Config
	ctx         context.Context
	mockRepo    *mocks.BookRepository
	bookService services.BookService
}

func (s *BookServiceSuite) SetupTest() {
	// Initialize context
	s.ctx = context.Background()

	// Initialize configuration
	s.cfg = &configs.Config{}

	// Initialize mock repository
	s.mockRepo = new(mocks.BookRepository)

	// Initialize book service
	s.bookService = services.NewBookService(s.cfg, s.mockRepo)
}

func TestBookService(t *testing.T) {
	suite.Run(t, new(BookServiceSuite))
}

func (s *BookServiceSuite) TestCreateNewBookServiceWhenSuccess() {
	dummyRequest := request.CreateBookRequest{
		Title:             "Test Book",
		Author:            "Test Author",
		Description:       "Test Description",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
	}

	s.mockRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*entities.Book")).Return(nil)

	err := s.bookService.Create(s.ctx, &dummyRequest)
	s.NoError(err)
}

func (s *BookServiceSuite) TestCreateNewBookServiceWhenFailedValidation() {
	dummyRequest := request.CreateBookRequest{
		Title:             "",
		Author:            "Test Author",
		Description:       "Test Description",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
	}

	s.mockRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*entities.Book")).Return(nil)

	err := s.bookService.Create(s.ctx, &dummyRequest)
	s.Error(err)
}

func (s *BookServiceSuite) TestCreateNewBookServiceWhenFailedConnection() {
	dummyRequest := request.CreateBookRequest{
		Title:             "Test Book",
		Author:            "Test Author",
		Description:       "Test Description",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
	}

	s.mockRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*entities.Book")).Return(sql.ErrConnDone)

	err := s.bookService.Create(s.ctx, &dummyRequest)
	s.Error(err)
}

func (s *BookServiceSuite) TestFindAllBookServiceWhenSuccess() {
	updatedAt := time.Now()

	// Dummy data
	dummyBooks := []*entities.Book{
		{
			ID:                "1",
			ISBN:              "123-456-789",
			Title:             "Test Book 1",
			Description:       "Description 1",
			Author:            "Author 1",
			YearOfManufacture: 2023,
			Stock:             10,
			Price:             10000,
			IsDisplayed:       true,
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
		},
		{
			ID:                "2",
			ISBN:              "123-456-789",
			Title:             "Test Book 2",
			Description:       "Description 2",
			Author:            "Author 2",
			YearOfManufacture: 2023,
			Stock:             10,
			Price:             10000,
			IsDisplayed:       true,
			CreatedAt:         time.Now(),
			UpdatedAt:         &updatedAt,
		},
	}

	// Setup mock expectations
	s.mockRepo.On("FindAll", mock.Anything, 1, 10).Return(dummyBooks, int64(len(dummyBooks)), nil)

	// Call the service method
	books, totalRecords, err := s.bookService.FindAll(context.Background(), 1, 10)

	require.NoError(s.T(), err)
	require.Equal(s.T(), len(dummyBooks), len(books))
	require.Equal(s.T(), int(len(dummyBooks)), totalRecords)
}

func (s *BookServiceSuite) TestFindAllBookServiceWhenFailedConnection() {
	// Setup mock expectations
	s.mockRepo.On("FindAll", mock.Anything, 1, 10).Return(nil, int64(0), sql.ErrConnDone)

	// Call the service method
	books, totalRecords, err := s.bookService.FindAll(context.Background(), 1, 10)

	require.Error(s.T(), err)
	require.Nil(s.T(), books)
	require.Equal(s.T(), int(0), totalRecords)
}

func (s *BookServiceSuite) TestFindAllSearchBooksServiceWhenSuccess() {
	// Dummy data
	dummyBooks := []*entities.Book{
		{
			ID:                "1",
			ISBN:              "123-456-789",
			Title:             "Test Book 1",
			Description:       "Description 1",
			Author:            "Author 1",
			YearOfManufacture: 2023,
			Stock:             10,
			Price:             10000,
			IsDisplayed:       true,
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
		},
		{
			ID:                "2",
			ISBN:              "123-456-789",
			Title:             "Test Book 2",
			Description:       "Description 2",
			Author:            "Author 2",
			YearOfManufacture: 2023,
			Stock:             10,
			Price:             10000,
			IsDisplayed:       true,
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
		},
	}

	// dummy search
	dummySearch := &request.SearchBookRequest{
		Title:  "Test Book",
		Author: "Test Author",
		Price:  10000,
	}

	// Setup mock expectations
	s.mockRepo.On("FindAllSearch", mock.Anything, 1, 10, mock.AnythingOfType("*request.SearchBookRequest")).Return(dummyBooks, int64(len(dummyBooks)), nil)

	// Call the service method
	books, totalRecords, err := s.bookService.FindAllSearch(context.Background(), 1, 10, dummySearch)

	require.NoError(s.T(), err)
	require.Equal(s.T(), len(dummyBooks), len(books))
	require.Equal(s.T(), int(len(dummyBooks)), totalRecords)
}

func (s *BookServiceSuite) TestFindAllSearchBooksServiceWhenFailedConnection() {
	// dummy search
	dummySearch := &request.SearchBookRequest{
		Title:  "Test Book",
		Author: "Test Author",
		Price:  10000,
	}

	// Setup mock expectations
	s.mockRepo.On("FindAllSearch", mock.Anything, 1, 10, mock.AnythingOfType("*request.SearchBookRequest")).Return(nil, int64(0), sql.ErrConnDone)

	// Call the service method
	books, totalRecords, err := s.bookService.FindAllSearch(context.Background(), 1, 10, dummySearch)

	require.Error(s.T(), err)
	require.Nil(s.T(), books)
	require.Equal(s.T(), int(0), totalRecords)
}

func (s *BookServiceSuite) TestFindByIDBookServiceWhenSuccess() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)

	// Call the service method
	book, err := s.bookService.FindByID(s.ctx, "1")

	require.NoError(s.T(), err)
	require.Equal(s.T(), dummyBook.ID, book.Id)
	require.Equal(s.T(), dummyBook.ISBN, book.Isbn)
	require.Equal(s.T(), dummyBook.Title, book.Title)
	require.Equal(s.T(), dummyBook.Author, book.Author)
	require.Equal(s.T(), dummyBook.YearOfManufacture, book.YearOfManufacture)
	require.Equal(s.T(), dummyBook.Stock, book.Stock)
	require.Equal(s.T(), dummyBook.Price, book.Price)
	require.Equal(s.T(), dummyBook.IsDisplayed, book.IsDisplayed)
	require.Equal(s.T(), dummyBook.CreatedAt.String(), book.CreatedAt)
	if dummyBook.UpdatedAt != nil {
		require.Equal(s.T(), dummyBook.UpdatedAt.String(), book.UpdatedAt)
	} else {
		require.Empty(s.T(), book.UpdatedAt)
	}
}

func (s *BookServiceSuite) TestFindByIDBookServiceWhenFailedNotFound() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, sql.ErrNoRows)

	// Call the service method
	book, err := s.bookService.FindByID(s.ctx, "1")

	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookServiceSuite) TestFindByIDBookServiceWhenFailedConnection() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, sql.ErrConnDone)

	// Call the service method
	book, err := s.bookService.FindByID(s.ctx, "1")

	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookServiceSuite) TestFindByISBNBookServiceWhenSuccess() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByISBN", s.ctx, "123-456-789").Return(dummyBook, nil)

	// Call the service method
	book, err := s.bookService.FindByISBN(s.ctx, "123-456-789")

	require.NoError(s.T(), err)
	require.Equal(s.T(), dummyBook.ID, book.Id)
	require.Equal(s.T(), dummyBook.ISBN, book.Isbn)
	require.Equal(s.T(), dummyBook.Title, book.Title)
	require.Equal(s.T(), dummyBook.Author, book.Author)
	require.Equal(s.T(), dummyBook.YearOfManufacture, book.YearOfManufacture)
	require.Equal(s.T(), dummyBook.Stock, book.Stock)
	require.Equal(s.T(), dummyBook.Price, book.Price)
	require.Equal(s.T(), dummyBook.IsDisplayed, book.IsDisplayed)
	require.Equal(s.T(), dummyBook.CreatedAt.String(), book.CreatedAt)
	if dummyBook.UpdatedAt != nil {
		require.Equal(s.T(), dummyBook.UpdatedAt.String(), book.UpdatedAt)
	} else {
		require.Empty(s.T(), book.UpdatedAt)
	}
}

func (s *BookServiceSuite) TestFindByISBNBookServiceWhenFailedNotFound() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByISBN", s.ctx, "123-456-789").Return(dummyBook, sql.ErrNoRows)

	// Call the service method
	book, err := s.bookService.FindByISBN(s.ctx, "123-456-789")

	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookServiceSuite) TestFindByISBNBookServiceWhenFailedConnection() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByISBN", s.ctx, "123-456-789").Return(dummyBook, sql.ErrConnDone)

	// Call the service method
	book, err := s.bookService.FindByISBN(s.ctx, "123-456-789")

	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookServiceSuite) TestUpdateBookServiceWhenSuccess() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// dummy book request
	dummyRequest := request.UpdateBookRequest{
		Title:             "Test Book 1 Updated",
		Description:       "Description 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
	}

	updateAt := time.Now().Local()
	dummyBook.UpdatedAt = &updateAt

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("Update", s.ctx, dummyBook).Return(nil)

	// Call the service method
	err := s.bookService.Update(s.ctx, "1", &dummyRequest)

	require.NoError(s.T(), err)
}

func (s *BookServiceSuite) TestUpdateBookServiceWhenFailedNotFound() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// dummy book request
	dummyRequest := request.UpdateBookRequest{
		Title:             "Test Book 1 Updated",
		Description:       "Description 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
	}

	updatedAt := time.Now().Local()
	dummyBook.UpdatedAt = &updatedAt

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(nil, sql.ErrNoRows)
	s.mockRepo.On("Update", s.ctx, dummyBook).Return(nil)

	// Call the service method
	err := s.bookService.Update(s.ctx, "1", &dummyRequest)

	require.Error(s.T(), err)
}

func (s *BookServiceSuite) TestUpdateBookServiceWhenFailedValidation() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// dummy book request
	dummyRequest := request.UpdateBookRequest{
		Title:             "",
		Description:       "Description 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
	}

	updatedAt := time.Now().Local()
	dummyBook.UpdatedAt = &updatedAt

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("Update", s.ctx, dummyBook).Return(nil)

	// Call the service method
	err := s.bookService.Update(s.ctx, "1", &dummyRequest)

	require.Error(s.T(), err)
}

func (s *BookServiceSuite) TestUpdateBookServiceWhenFailedConnection() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// dummy book request
	dummyRequest := request.UpdateBookRequest{
		Title:             "Test Book 1 Updated",
		Description:       "Description 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
	}

	updatedAt := time.Now().Local()
	dummyBook.UpdatedAt = &updatedAt

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("Update", s.ctx, dummyBook).Return(sql.ErrConnDone)

	// Call the service method
	err := s.bookService.Update(s.ctx, "1", &dummyRequest)

	require.Error(s.T(), err)
}

func (s *BookServiceSuite) TestDeleteBookByIDServiceWhenSuccess() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("DeleteByID", s.ctx, "1").Return(nil)

	// Call the service method
	err := s.bookService.DeleteByID(s.ctx, "1")

	require.NoError(s.T(), err)
}

func (s *BookServiceSuite) TestDeleteBookByIDServiceWhenFailedNotFound() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("DeleteByID", s.ctx, "1").Return(sql.ErrNoRows)

	// Call the service method
	err := s.bookService.DeleteByID(s.ctx, "1")

	require.Error(s.T(), err)
}

func (s *BookServiceSuite) TestDeleteBookByIDServiceWhenFailedConnection() {
	// Dummy data
	dummyBook := &entities.Book{
		ID:                "1",
		ISBN:              "123-456-789",
		Title:             "Test Book 1",
		Description:       "Description 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             10000,
		IsDisplayed:       true,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         nil,
	}

	// Setup mock expectations
	s.mockRepo.On("FindByID", s.ctx, "1").Return(dummyBook, nil)
	s.mockRepo.On("DeleteByID", s.ctx, "1").Return(sql.ErrConnDone)

	// Call the service method
	err := s.bookService.DeleteByID(s.ctx, "1")

	require.Error(s.T(), err)
}

func (s *BookServiceSuite) TestDeleteAllBooksServiceWhenSuccess() {
	// Setup mock expectations
	s.mockRepo.On("DeleteAll", s.ctx).Return(nil)

	// Call the service method
	err := s.bookService.DeleteAll(s.ctx)

	require.NoError(s.T(), err)
}

func (s *BookServiceSuite) TestDeleteAllBooksServiceWhenFailedConnection() {
	// Setup mock expectations
	s.mockRepo.On("DeleteAll", s.ctx).Return(sql.ErrConnDone)

	// Call the service method
	err := s.bookService.DeleteAll(s.ctx)

	require.Error(s.T(), err)
}
