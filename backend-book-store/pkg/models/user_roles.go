package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}