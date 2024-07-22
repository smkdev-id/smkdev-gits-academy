package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/helper"
	"github.com/smkdev-id/smkdev-gits-academy/tree/ahmadyani/backend/golang-todo-restful-api/model/domain"
)

type TodoListRepositoryImpl struct {
}

func NewTodoListRepository() TodoListRepository {
	return &TodoListRepositoryImpl{}
}

func (repository *TodoListRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todolist domain.TodoList) domain.TodoList {
	SQL := "insert into todolist (description, completed) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, todolist.Description, todolist.Completed)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todolist.Id = int(id)
	return todolist
}

func (repository *TodoListRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todolist domain.TodoList) domain.TodoList {
	SQL := "update todolist set description = ?, completed = ? where id =?"
	_, err := tx.ExecContext(ctx, SQL, todolist.Description, todolist.Completed, todolist.Id)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodoListRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todolist domain.TodoList) {
	SQL := "delete from todolist where id =?"
	_, err := tx.ExecContext(ctx, SQL, todolist.Id)
	helper.PanicIfError(err)
}

func (repository *TodoListRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todolistId int) (domain.TodoList, error) {
	SQL := "select id, description, completed from todolist where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, todolistId)
	helper.PanicIfError(err)
	defer rows.Close()

	todolist := domain.TodoList{}
	if rows.Next() {
		err := rows.Scan(&todolist.Id, &todolist.Description, &todolist.Completed)
		helper.PanicIfError(err)
		return todolist, nil
	} else {
		return todolist, errors.New("todolist not found")
	}

}

func (repository *TodoListRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList {
	SQL := "select id, description, completed from todolist"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todoLists []domain.TodoList
	for rows.Next() {
		todolist := domain.TodoList{}
		err := rows.Scan(&todolist.Id, &todolist.Description, &todolist.Completed)
		helper.PanicIfError(err)
		todoLists = append(todoLists, todolist)
	}

	return todoLists
}
