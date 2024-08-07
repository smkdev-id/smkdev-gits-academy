package repository

import (
	"EkoEdyPurwanto/task-3/internal/model"
	"EkoEdyPurwanto/task-3/internal/model/dto/req"
)

type (
	BookRepository interface {
		Save(book model.Book) error
		FindAll() (model.Books, error)
		FindById(id string) (model.Books, error)
		Modify(identifier req.UpdateRequest) error
		Delete(identifier req.DeleteRequest) error
	}
)
