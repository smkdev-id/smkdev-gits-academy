package model

import (
	"time"
)

type Status string
type Priority string

const (
	StatusPending     Status = "pending"
	StatusInProgress  Status = "in_progress"
	StatusCompleted   Status = "completed"

	PriorityLow       Priority = "low"
	PriorityMedium    Priority = "medium"
	PriorityHigh      Priority = "high"
)

type Todo struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Task      string    `json:"task" gorm:"type:varchar(255);not null"`
	DueDate   time.Time `json:"due_date" gorm:"type:date"`
	Status    Status    `json:"status" gorm:"type:enum('pending', 'in_progress', 'completed');default:'pending'"`
	Priority  Priority  `json:"priority" gorm:"type:enum('low', 'medium', 'high');default:'medium'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
