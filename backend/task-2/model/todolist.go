package model

import (
	"database/sql"
)

type (
	Status string

	TodoList struct {
		Id          string
		Title       string
		Description string
		Status      Status
		CreatedAt   sql.NullTime
		UpdatedAt   sql.NullTime
		DeletedAt   sql.NullTime
	}

	TodoLists []TodoList
)

const (
	StatusPending  Status = "pending"
	StatusComplete Status = "completed"
	// other statuses
)
