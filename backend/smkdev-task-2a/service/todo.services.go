package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
}

func (t *Todo) Create(db *sql.DB) error {
	t.Id = uuid.New().String()
	t.CreatedAt = time.Now()
	_, err := db.Exec("INSERT INTO todos (id, createdAt, title, description, status) VALUES (?, ?, ?, ?, ?)", t.Id, t.CreatedAt, t.Title, t.Description, t.Status)
	if err != nil {
		fmt.Println("Error inserting todo:", err)
	}
	return err
}

func GetTodoById(db *sql.DB, Id string) (*Todo, error) {
	t := new(Todo)
	err := db.QueryRow("SELECT id, createdAt, title, description, status FROM todos WHERE id = ?", Id).Scan(&t.Id, &t.CreatedAt, &t.Title, &t.Description, &t.Status)
	if err != nil {
		fmt.Println("Error fetching todo by id:", err)
		return nil, err
	}
	return t, nil
}

func (t *Todo) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE todos SET title = ?, description = ?, status = ? WHERE id = ?", t.Title, t.Description, t.Status, t.Id)
	if err != nil {
		fmt.Println("Error updating todo:", err)
	}
	return err
}

func DeleteTodoById(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error deleting todo:", err)
	}
	return err
}

func GetAllTodos(db *sql.DB) ([]*Todo, error) {
	rows, err := db.Query("SELECT id, createdAt, title, description, status FROM todos")
	if err != nil {
		fmt.Println("Error fetching all todos:", err)
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		t := new(Todo)
		err := rows.Scan(&t.Id, &t.CreatedAt, &t.Title, &t.Description, &t.Status)
		if err != nil {
			fmt.Println("Error scanning todo row:", err)
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}
