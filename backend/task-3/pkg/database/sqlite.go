package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alwiirfan/configs"
	"github.com/alwiirfan/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SqliteConnection(ctx context.Context, cfg *configs.Config) (*gorm.DB, error) {

	var log logger.Interface

	if cfg.AppEnv == "development" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default.LogMode(logger.Silent)
	}

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		return nil, errors.New("DATABASE_PATH is not set")
	}
	fmt.Printf("Connecting to database at path: %s", dbPath)

	// open sqlite database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set database connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping database to check connection
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// if database is not ready, return error
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	DB = db

	// migrate table to database
	err = DB.AutoMigrate(&entities.Book{})
	if err != nil {
		return nil, err
	}

	return db, err
}
