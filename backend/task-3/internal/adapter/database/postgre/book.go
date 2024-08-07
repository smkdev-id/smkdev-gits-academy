package postgre

import (
	"EkoEdyPurwanto/task-3/internal/model"
	"EkoEdyPurwanto/task-3/internal/model/dto/req"
	"EkoEdyPurwanto/task-3/internal/repository"
	"database/sql"
)

type bookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) repository.BookRepository {
	return &bookRepository{
		DB: db,
	}
}

// Save implements repository.BookRepository.
func (b *bookRepository) Save(book model.Book) error {
	SQL := `INSERT INTO "book"(id, title, isbn, author, published_date, genre, price, quantity, created_at, updated_at, deleted_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := b.DB.Exec(SQL,
		&book.Id,
		&book.Title,
		&book.ISBN,
		&book.Author,
		&book.PublishedDate,
		&book.Genre,
		&book.Price,
		&book.Quantity,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements repository.BookRepository.
func (b *bookRepository) FindAll() (model.Books, error) {
	SQL := `SELECT id, title, isbn, author, published_date, genre, price, quantity, created_at, updated_at, deleted_at FROM "book"`

	rows, err := b.DB.Query(SQL)
	if err != nil {
		return nil, err
	}
	var books model.Books
	for rows.Next() {
		var book model.Book
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.ISBN,
			&book.Author,
			&book.PublishedDate,
			&book.Genre,
			&book.Price,
			&book.Quantity,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// FindById implements repository.BookRepository.
func (b *bookRepository) FindById(id string) (model.Books, error) {
	SQL := `SELECT id, title, isbn, author, published_date, genre, price, quantity, created_at, updated_at, deleted_at FROM "book" WHERE id = $1`

	rows, err := b.DB.Query(SQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books model.Books
	for rows.Next() {
		var book model.Book
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.ISBN,
			&book.Author,
			&book.PublishedDate,
			&book.Genre,
			&book.Price,
			&book.Quantity,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// Modify implements repository.BookRepository.
func (b *bookRepository) Modify(identifier req.UpdateRequest) error {
	SQL := `UPDATE "book" SET title = $2, isbn = $3, author = $4, published_date = 5, genre = $6, price = $7, quantity = $8, updated_at = $9 WHERE id = $1`

	_, err := b.DB.Exec(SQL,
		identifier.Id,
		identifier.Title,
		identifier.ISBN,
		identifier.Author,
		identifier.PublishedDate,
		identifier.Genre,
		identifier.Price,
		identifier.Quantity,
		identifier.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements repository.BookRepository.
func (b *bookRepository) Delete(identifier req.DeleteRequest) error {
	SQL := `DELETE FROM "book" WHERE id = $1`

	_, err := b.DB.Exec(SQL, identifier.Id)
	if err != nil {
		return err
	}

	return nil
}
