package model

import (
	"time"
)

type (
	Status string

	TodoList struct {
		Id          string
		Title       string
		Description string
		Status      Status
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   time.Time
	}

	TodoLists []TodoList
)

const (
	StatusPending  Status = "pending"
	StatusComplete Status = "completed"
	// other statuses
)
