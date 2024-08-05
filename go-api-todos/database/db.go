package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func init() {
	db := OpenDB()

	defer CloseDB(db)

	err := db.AutoMigrate(&Todo{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Database initialized")
}
func OpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}

	dbSQL.Close()
}

func GetTodos() []Todo {
	db := OpenDB()
	defer CloseDB(db)

	var todos []Todo
	err := db.Find(&todos).Error
	if err != nil {
		fmt.Println(err)
	}
	return todos
}

func CreateTodo(todo *Todo) {
	db := OpenDB()
	defer CloseDB(db)

	todo.ID = GenerateID(db)
	db.Create(&todo)
}

func GenerateID(db *gorm.DB) int {
	var lastTodo Todo
	err := db.Last(&Todo{}).Scan(&lastTodo).Error

	if err != nil {
		return 1
	}

	return lastTodo.ID + 1
}

func UpdateTodo(id string, todo *Todo) error {
	db := OpenDB()
	defer CloseDB(db)

	err := db.Model(&Todo{}).Where("id = ?", id).Updates(todo).Error
	return err
}

func DeleteTodo(id string) error {
	db := OpenDB()
	defer CloseDB(db)

	err := db.Where("id = ?", id).Delete(&Todo{}).Error

	return err
}
