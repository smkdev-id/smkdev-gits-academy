package services

import (
	"database/sql"
	"strings"
	"todolist/config"
	"todolist/models"
)

// GetTodos mengambil semua todo dari database
func GetTodos() ([]models.Todo, error) {
	// Menjalankan query untuk mengambil semua todo
	rows, err := config.DB.Query("SELECT id, todo, completed FROM todolist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	// Iterasi melalui hasil query dan masukkan ke dalam slice todos
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Todo, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// CreateTodo membuat todo baru di database
func CreateTodo(todo models.Todo) (models.Todo, error) {
	result, err := config.DB.Exec("INSERT INTO todolist (todo, completed) VALUES (?, ?)", todo.Todo, todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}
	// Mendapatkan ID dari todo yang baru dibuat
	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, err
	}

	todo.ID = int(id)

	return todo, nil
}

// UpdateTodo memperbarui todo yang ada di database berdasarkan ID
func UpdateTodo(id int, todo models.Todo) (models.Todo, error) {
	var existingTodo models.Todo
	// Mengambil todo yang ada dari database berdasarkan ID
	err := config.DB.QueryRow("SELECT id, todo, completed FROM todolist WHERE id = ?", id).Scan(&existingTodo.ID, &existingTodo.Todo, &existingTodo.Completed)
	if err == sql.ErrNoRows {
		return models.Todo{}, sql.ErrNoRows
	} else if err != nil {
		return models.Todo{}, err
	}

	updateFields := []string{}
	updateValues := []interface{}{}
	// Memeriksa apakah field 'todo' perlu diperbarui
	if todo.Todo != "" && todo.Todo != existingTodo.Todo {
		updateFields = append(updateFields, "todo = ?")
		updateValues = append(updateValues, todo.Todo)
	}
	// Memeriksa apakah field 'completed' perlu diperbarui
	if todo.Completed != existingTodo.Completed {
		updateFields = append(updateFields, "completed = ?")
		updateValues = append(updateValues, todo.Completed)
	}
	// Jika tidak ada field yang diperbarui, kembalikan existingTodo
	if len(updateFields) == 0 {
		return existingTodo, nil
	}

	updateValues = append(updateValues, id)
	// Membuat query update dinamis berdasarkan field yang perlu diperbarui
	updateQuery := "UPDATE todolist SET " + strings.Join(updateFields, ", ") + " WHERE id = ?"
	_, err = config.DB.Exec(updateQuery, updateValues...)
	if err != nil {
		return models.Todo{}, err
	}
	// Mengambil kembali todo yang diperbarui dari database
	err = config.DB.QueryRow("SELECT id, todo, completed FROM todolist WHERE id = ?", id).Scan(&todo.ID, &todo.Todo, &todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// DeleteTodo menghapus todo dari database berdasarkan ID
func DeleteTodo(id int) error {
	result, err := config.DB.Exec("DELETE FROM todolist WHERE id = ?", id)
	if err != nil {
		return err
	}
	// Memeriksa apakah ada baris yang terpengaruh oleh query delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
