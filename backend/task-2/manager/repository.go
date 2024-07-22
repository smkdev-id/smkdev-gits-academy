package manager

import "EkoEdyPurwanto/task-2/repository/postgres"

type (
	RepositoryManager interface {
		TodoListRepository() postgres.TodoListRepository
	}
	repositoryManager struct {
		InfraM InfraManager
	}
)

func (r *repositoryManager) TodoListRepository() postgres.TodoListRepository {
	return postgres.NewTodoListRepository(r.InfraM.Conn())
}

// NewRepositoryManager Constructor
func NewRepositoryManager(infraM InfraManager) RepositoryManager {
	return &repositoryManager{
		InfraM: infraM,
	}
}
