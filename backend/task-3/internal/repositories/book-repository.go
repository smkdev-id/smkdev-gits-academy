package repositories

import (
	"context"

	"github.com/alwiirfan/internal/dto/request"
	"github.com/alwiirfan/internal/entities"
	"github.com/alwiirfan/internal/utils/pagination"
	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAll(ctx context.Context, page, pageSize int) ([]*entities.Book, int64, error)
		FindAllSearch(ctx context.Context, page, pageSize int, req *request.SearchBookRequest) ([]*entities.Book, int64, error)
		FindByID(ctx context.Context, id string) (*entities.Book, error)
		FindByISBN(ctx context.Context, isbn string) (*entities.Book, error)
		Create(ctx context.Context, book *entities.Book) error
		Update(ctx context.Context, book *entities.Book) error
		DeleteByID(ctx context.Context, id string) error
		DeleteAll(ctx context.Context) error
	}

	bookRepository struct {
		DB *gorm.DB
	}
)

// Create implements BookRepository.
func (r *bookRepository) Create(ctx context.Context, book *entities.Book) error {
	if err := r.DB.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements BookRepository.
func (r *bookRepository) DeleteByID(ctx context.Context, id string) error {
	var book *entities.Book

	if err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAll implements BookRepository.
func (r *bookRepository) DeleteAll(ctx context.Context) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM books").Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll implements BookRepository.
func (r *bookRepository) FindAll(ctx context.Context, page, pageSize int) ([]*entities.Book, int64, error) {
	var books []*entities.Book
	var totalRecords int64

	if err := r.DB.WithContext(ctx).Model(&entities.Book{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.WithContext(ctx).Scopes(pagination.BookPaginate(page, pageSize)).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, totalRecords, nil
}

// Fina All Books with optional search params
func (r *bookRepository) FindAllSearch(ctx context.Context, page, pageSize int, req *request.SearchBookRequest) ([]*entities.Book, int64, error) {
	var books []*entities.Book
	var totalRecords int64

	query := r.DB.WithContext(ctx).Model(&entities.Book{})

	if req.ISBN != "" {
		query = query.Where("isbn LIKE ?", "%"+req.ISBN+"%")
	}

	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}

	if req.Year != 0 {
		query = query.Where("year_of_manufacture = ?", req.Year)
	}

	if req.Author != "" {
		query = query.Where("author LIKE ?", "%"+req.Author+"%")
	}

	if req.Price != 0 {
		query = query.Where("price = ?", req.Price)
	}

	if req.IsDisplayed == nil {
		req.IsDisplayed = new(bool)
		*req.IsDisplayed = true
	}

	if req.StartDate != "" && req.EndDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	} else if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	} else if req.EndDate != "" {
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.WithContext(ctx).Scopes(pagination.BookPaginate(page, pageSize)).Where("is_displayed = ?", &req.IsDisplayed).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, totalRecords, nil
}

// FindByID implements BookRepository.
func (r *bookRepository) FindByID(ctx context.Context, id string) (*entities.Book, error) {
	var book *entities.Book

	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// FindByISBN implements BookRepository.
func (r *bookRepository) FindByISBN(ctx context.Context, isbn string) (*entities.Book, error) {
	var book *entities.Book

	if err := r.DB.WithContext(ctx).Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// Update implements BookRepository.
func (r *bookRepository) Update(ctx context.Context, book *entities.Book) error {
	if err := r.DB.WithContext(ctx).Model(&entities.Book{}).Where("id = ?", book.ID).Updates(book).Error; err != nil {
		return err
	}

	return nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}
