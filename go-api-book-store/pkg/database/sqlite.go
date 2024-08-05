package database

import (
	"github.com/hasban-fardani/go-api-book-store/pkg/config"
	models "github.com/hasban-fardani/go-api-book-store/pkg/model"
	"github.com/hasban-fardani/go-api-book-store/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	db := OpenDB()
	db.AutoMigrate(&models.Book{})

	CloseDB(db)
}

func OpenDB() *gorm.DB {
	// load config from .env file
	cfg, err := config.NewConfig(".env")

	utils.HandleError(err)

	var log logger.Interface
	if cfg.Env == "dev" {
		log = logger.Default
	} else {
		log = logger.Discard
	}

	db, err := gorm.Open(sqlite.Open(cfg.DBName), &gorm.Config{
		Logger: log,
	})

	utils.HandleError(err)

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	utils.HandleError(err)
	dbSQL.Close()
}
