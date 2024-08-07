package manager

import (
	"EkoEdyPurwanto/task-3/internal/adapter/database/postgre"
	"EkoEdyPurwanto/task-3/internal/repository"
)

type (
	RepositoryManager interface {
		BookRepository() repository.BookRepository
	}
	repositoryManager struct {
		InfraM InfraManager
	}
)

func (r *repositoryManager) BookRepository() repository.BookRepository {
	return postgre.NewBookRepository(r.InfraM.Conn())
}

// NewRepositoryManager Constructor
func NewRepositoryManager(infraM InfraManager) RepositoryManager {
	return &repositoryManager{
		InfraM: infraM,
	}
}
