package config

import (
	"context"
	"log"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/models"
	"github.com/go-redis/redis/v8"
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
	DB.AutoMigrate(&models.Book{}, &models.User{}, &models.Role{})
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