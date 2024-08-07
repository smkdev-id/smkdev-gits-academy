package config

import (
	"context"
	"log"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/models"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Role{}, &models.UserRole{})
	
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to redis: ", err)
	}
}

func SeedData() {
	adminRole := models.Role{Name: "admin"}
	userRole := models.Role{Name: "user"}

	if err := DB.FirstOrCreate(&adminRole, models.Role{Name: "admin"}).Error; err != nil {
		log.Fatalf("Error creating admin role: %v", err)
	}
	if err := DB.FirstOrCreate(&userRole, models.Role{Name: "user"}).Error; err != nil {
		log.Fatalf("Error creating user role: %v", err)
	}

	admin := models.User{Username: "admin", Password: "password123"}
	user := models.User{Username: "user", Password: "password123"}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	admin.Password = string(hashedPassword)
	user.Password = string(hashedPassword)

	if err := DB.FirstOrCreate(&admin, models.User{Username: "admin"}).Error; err != nil {
		log.Fatalf("Error creating admin user: %v", err)
	}
	if err := DB.FirstOrCreate(&user, models.User{Username: "user"}).Error; err != nil {
		log.Fatalf("Error creating user user: %v", err)
	}

	if err := DB.Model(&admin).Association("Roles").Append(&adminRole); err != nil {
		log.Fatalf("Error assigning role to admin: %v", err)
	}
	if err := DB.Model(&user).Association("Roles").Append(&userRole); err != nil {
		log.Fatalf("Error assigning role to user: %v", err)
	}
}