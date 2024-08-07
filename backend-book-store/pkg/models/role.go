package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name", gorm:"unique"`
	Users       []User       `gorm:"many2many:user_roles"`
}