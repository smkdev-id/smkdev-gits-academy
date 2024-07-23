package postgres

import (
	"EkoEdyPurwanto/task-2/model"
	"EkoEdyPurwanto/task-2/model/dto/req"
	"database/sql"
)

type (
	TodoListRepository interface {
		Save(todoList model.TodoList) error
		FindAll() (model.TodoLists, error)
		FindByIdentifier(identifier string) (model.TodoLists, error)
		Update(identifier req.UpdateRequest) error
		Delete(id string) error
	}

	todoListRepository struct {
		DB *sql.DB
	}
)

// Save implements TodoListRepository.
func (t *todoListRepository) Save(todoList model.TodoList) error {
	SQL := `INSERT INTO "todo"(id, title, description, status, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := t.DB.Exec(SQL,
		&todoList.Id,
		&todoList.Title,
		&todoList.Description,
		&todoList.Status,
		&todoList.CreatedAt,
		&todoList.UpdatedAt,
		&todoList.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements TodoListRepository.
func (t *todoListRepository) FindAll() (model.TodoLists, error) {
	SQL := `SELECT id, title, description, status, created_at, updated_at, deleted_at FROM "todo"`

	rows, err := t.DB.Query(SQL)
	if err != nil {
		return nil, err
	}
	var todolists model.TodoLists
	for rows.Next() {
		var todolist model.TodoList
		err := rows.Scan(
			&todolist.Id,
			&todolist.Title,
			&todolist.Description,
			&todolist.Status,
			&todolist.CreatedAt,
			&todolist.UpdatedAt,
			&todolist.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		todolists = append(todolists, todolist)
	}
	return todolists, nil
}

// FindByIdentifier implements TodoListRepository.
func (t *todoListRepository) FindByIdentifier(identifier string) (model.TodoLists, error) {
	SQL := `SELECT id, title, description, status, created_at, updated_at, deleted_at FROM "todo" WHERE status = $1 `

	rows, err := t.DB.Query(SQL, identifier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todolists model.TodoLists
	for rows.Next() {
		var todolist model.TodoList
		err := rows.Scan(
			&todolist.Id,
			&todolist.Title,
			&todolist.Description,
			&todolist.Status,
			&todolist.CreatedAt,
			&todolist.UpdatedAt,
			&todolist.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		todolists = append(todolists, todolist)
	}
	return todolists, nil
}

// Update implements TodoListRepository.
func (t *todoListRepository) Update(identifier req.UpdateRequest) error {
	SQL := `UPDATE "todo" SET status = $2, updated_at = $3 WHERE id = $1`

	_, err := t.DB.Exec(SQL, identifier.Id, identifier.Status, identifier.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements TodoListRepository.
func (t *todoListRepository) Delete(id string) error {
	SQL := `DELETE FROM "todo" WHERE id = $1`

	_, err := t.DB.Exec(SQL, id)
	if err != nil {
		return err
	}

	return nil
}

// NewTodoListRepository Constructor
func NewTodoListRepository(db *sql.DB) TodoListRepository {
	return &todoListRepository{
		DB: db,
	}
}
