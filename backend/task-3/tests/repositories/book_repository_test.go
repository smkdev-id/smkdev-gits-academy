package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/entities"
	"github.com/alwiirfan/internal/repositories"
	"github.com/alwiirfan/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BookRepositorySuite struct {
	suite.Suite
	DB *gorm.DB

	book       *entities.Book
	repository repositories.BookRepository
}

// SetupSuite initializes the database
func (s *BookRepositorySuite) SetupSuite() {
	var err error

	// Initialize database
	s.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(s.T(), err)

	// Migrate the database
	err = s.DB.AutoMigrate(s.book)
	require.NoError(s.T(), err)

	// Initialize repository
	s.repository = repositories.NewBookRepository(s.DB)
}

// TearDownSuite clears the database after each test
func (s *BookRepositorySuite) TearDownSuite() {
	sqlDB, err := s.DB.DB()
	require.NoError(s.T(), err)
	sqlDB.Close()
}

// BeforeTest clears the database before each test
func (s *BookRepositorySuite) BeforeTest(suiteName, testName string) {
	s.DB.Exec("DELETE FROM books")
}

func TestBookRepository(t *testing.T) {
	suite.Run(t, new(BookRepositorySuite))
}

func (s *BookRepositorySuite) TestCreateNewBookRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book",
		Description:       "This is a test book",
		Author:            "John Doe",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}

	err := s.repository.Create(context.Background(), book)
	require.NoError(s.T(), err)

	var foundBook entities.Book
	err = s.DB.Where("id = ?", book.ID).First(&foundBook).Error
	require.NoError(s.T(), err)

	require.Equal(s.T(), book.Title, foundBook.Title)
	require.Equal(s.T(), book.Description, foundBook.Description)
	require.Equal(s.T(), book.Author, foundBook.Author)
	require.Equal(s.T(), book.YearOfManufacture, foundBook.YearOfManufacture)
	require.Equal(s.T(), book.Stock, foundBook.Stock)
	require.Equal(s.T(), book.Price, foundBook.Price)
	require.Equal(s.T(), book.IsDisplayed, foundBook.IsDisplayed)
	require.WithinDuration(s.T(), book.CreatedAt, foundBook.CreatedAt, time.Second)
	require.Equal(s.T(), book.UpdatedAt, foundBook.UpdatedAt, time.Second)
}

func (s *BookRepositorySuite) TestCreateNewBookRepositoryWhenFailedConnection() {
	// Create a new repository with a disconnected DB
	badDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(s.T(), err)

	sqlDB, err := badDB.DB()
	require.NoError(s.T(), err)
	sqlDB.Close()

	// Try to create a new book with a closed connection
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Disconnected Book",
		Description:       "This is a disconnected book",
		Author:            "Jane Doe",
		YearOfManufacture: 2024,
		Stock:             5,
		Price:             200,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}

	disconnectedRepo := repositories.NewBookRepository(badDB)
	err = disconnectedRepo.Create(context.Background(), book)
	require.Error(s.T(), err)
}

func (s *BookRepositorySuite) TestFindAllBooksRepositoryWhenSuccess() {
	// Add sample data
	book1 := &entities.Book{ID: "1", ISBN: "123-456-789", Title: "Test Book 1", Author: "Author 1"}
	book2 := &entities.Book{ID: "2", ISBN: "987-654-321", Title: "Test Book 2", Author: "Author 2"}
	s.DB.Create(book1)
	s.DB.Create(book2)

	books, total, err := s.repository.FindAll(context.Background(), 1, 10)
	require.NoError(s.T(), err)
	require.Equal(s.T(), int64(2), total)
	require.Len(s.T(), books, 2)
}

func (s *BookRepositorySuite) TestFindAllBooksRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	books, total, err := repo.FindAll(context.Background(), 1, 10)
	require.Error(s.T(), err)
	require.Nil(s.T(), books)
	require.Equal(s.T(), int64(0), total)
}

func (s *BookRepositorySuite) TestFindByIDBookRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil}
	s.DB.Create(book)

	foundBook, err := s.repository.FindByID(context.Background(), book.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), book.ID, foundBook.ID)
	require.Equal(s.T(), book.ISBN, foundBook.ISBN)
	require.Equal(s.T(), book.Title, foundBook.Title)
	require.Equal(s.T(), book.Author, foundBook.Author)
	require.Equal(s.T(), book.YearOfManufacture, foundBook.YearOfManufacture)
	require.Equal(s.T(), book.Stock, foundBook.Stock)
	require.Equal(s.T(), book.Price, foundBook.Price)
	require.Equal(s.T(), book.IsDisplayed, foundBook.IsDisplayed)
	require.WithinDuration(s.T(), book.CreatedAt, foundBook.CreatedAt, time.Second)
	require.Equal(s.T(), book.UpdatedAt, foundBook.UpdatedAt)
}

func (s *BookRepositorySuite) TestFindByIDBookRepositoryWhenFailedNotFound() {
	book, err := s.repository.FindByID(context.Background(), "id-non-existent")
	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookRepositorySuite) TestFindByIDBookRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	foundBook, err := repo.FindByID(context.Background(), utils.GenerateUUID())
	require.Error(s.T(), err)
	require.Nil(s.T(), foundBook)
}

func (s *BookRepositorySuite) TestFindByISBNBookRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	s.DB.Create(book)

	foundBook, err := s.repository.FindByISBN(context.Background(), book.ISBN)
	require.NoError(s.T(), err)
	require.Equal(s.T(), book.ID, foundBook.ID)
	require.Equal(s.T(), book.ISBN, foundBook.ISBN)
	require.Equal(s.T(), book.Title, foundBook.Title)
	require.Equal(s.T(), book.Author, foundBook.Author)
	require.Equal(s.T(), book.YearOfManufacture, foundBook.YearOfManufacture)
	require.Equal(s.T(), book.Stock, foundBook.Stock)
	require.Equal(s.T(), book.Price, foundBook.Price)
	require.Equal(s.T(), book.IsDisplayed, foundBook.IsDisplayed)
	require.WithinDuration(s.T(), book.CreatedAt, foundBook.CreatedAt, time.Second)
	require.Equal(s.T(), book.UpdatedAt, foundBook.UpdatedAt)
}

func (s *BookRepositorySuite) TestFindByISBNBookRepositoryWhenFailedNotFound() {
	book, err := s.repository.FindByISBN(context.Background(), "isbn-non-existent")
	require.Error(s.T(), err)
	require.Nil(s.T(), book)
}

func (s *BookRepositorySuite) TestFindByISBNBookRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	foundBook, err := repo.FindByISBN(context.Background(), utils.GenerateISBN())
	require.Error(s.T(), err)
	require.Nil(s.T(), foundBook)
}

func (s *BookRepositorySuite) TestFindAllSearchBooksRepositoryWhenSuccess() {
	book1 := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	book2 := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 2",
		Author:            "Author 2",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	err := s.repository.Create(context.Background(), book1)
	require.NoError(s.T(), err)
	err = s.repository.Create(context.Background(), book2)
	require.NoError(s.T(), err)

	// Perform search
	req := &request.SearchBookRequest{
		Title: "Test Book",
	}
	books, total, err := s.repository.FindAllSearch(context.Background(), 1, 10, req)
	require.NoError(s.T(), err)

	// Check if we have the expected number of items
	require.Len(s.T(), books, 2) // Adjust this based on your test data
	require.Equal(s.T(), int64(2), total)

	// Optional: Verify book details
	require.Equal(s.T(), "Test Book 1", books[0].Title)
	require.Equal(s.T(), "Test Book 2", books[1].Title)
}

func (s *BookRepositorySuite) TestFindAllSearchBooksRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	_, _, err = repo.FindAllSearch(context.Background(), 1, 10, &request.SearchBookRequest{})
	require.Error(s.T(), err)
}

func (s *BookRepositorySuite) TestUpateBookRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	s.DB.Create(book)

	updatedAtnew := time.Now()
	updatedBook := &entities.Book{
		ID:                book.ID,
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         book.CreatedAt,
		UpdatedAt:         &updatedAtnew,
	}
	s.DB.Save(updatedBook)

	foundBook, err := s.repository.FindByID(context.Background(), book.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), updatedBook.ID, foundBook.ID)
	require.Equal(s.T(), updatedBook.ISBN, foundBook.ISBN)
	require.Equal(s.T(), updatedBook.Title, foundBook.Title)
	require.Equal(s.T(), updatedBook.Author, foundBook.Author)
	require.Equal(s.T(), updatedBook.YearOfManufacture, foundBook.YearOfManufacture)
	require.Equal(s.T(), updatedBook.Stock, foundBook.Stock)
	require.Equal(s.T(), updatedBook.Price, foundBook.Price)
	require.Equal(s.T(), updatedBook.IsDisplayed, foundBook.IsDisplayed)
	require.WithinDuration(s.T(), updatedBook.CreatedAt, foundBook.CreatedAt, time.Second)
	if updatedBook.UpdatedAt != nil {
		require.NotNil(s.T(), foundBook.UpdatedAt)
		require.WithinDuration(s.T(), *updatedBook.UpdatedAt, *foundBook.UpdatedAt, time.Second)
	} else {
		require.Nil(s.T(), foundBook.UpdatedAt)
	}
}

func (s *BookRepositorySuite) TestUpateBookRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	updatedAtnew := time.Now()
	updatedBook := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1 Updated",
		Author:            "Author 1 Updated",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         &updatedAtnew,
	}
	err = repo.Update(context.Background(), updatedBook)
	require.Error(s.T(), err)
}

func (s *BookRepositorySuite) TestDeleteBookRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	s.DB.Create(book)

	err := s.repository.DeleteByID(context.Background(), book.ID)
	require.NoError(s.T(), err)
	bookFound, err := s.repository.FindByID(context.Background(), book.ID)
	require.Error(s.T(), err)
	require.Nil(s.T(), bookFound)
}

func (s *BookRepositorySuite) TestDeleteBookRepositoryWhenFailedConnection() {
	// Simulate a failed connection
	badDB, err := gorm.Open(sqlite.Open(""), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(s.T(), err)
	repo := repositories.NewBookRepository(badDB)

	err = repo.DeleteByID(context.Background(), utils.GenerateUUID())
	require.Error(s.T(), err)
}

func (s *BookRepositorySuite) TestDeleteAllBooksRepositoryWhenSuccess() {
	book := &entities.Book{
		ID:                utils.GenerateUUID(),
		ISBN:              utils.GenerateISBN(),
		Title:             "Test Book 1",
		Author:            "Author 1",
		YearOfManufacture: 2023,
		Stock:             10,
		Price:             100,
		IsDisplayed:       false,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}
	s.DB.Create(book)

	err := s.repository.DeleteAll(context.Background())
	require.NoError(s.T(), err)
	bookFound, err := s.repository.FindByID(context.Background(), book.ID)
	require.Error(s.T(), err)
	require.Nil(s.T(), bookFound)
}
