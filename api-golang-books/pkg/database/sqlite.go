package database

import (
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/config"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/models"
	"github.com/smkdev-id/smkdev-gits-academy/tree/rafi-cahyadi/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	db := ConnectToSQLite()
	db.AutoMigrate(&models.Book{})

	defer CloseDB(db)
}
func ConnectToSQLite() *gorm.DB {
	cfg, err := config.NewConfig(".env")

	utils.HandleError(err)

	var log logger.Interface

	if cfg.Env == "dev" {
		log = logger.Default
	} else {
		log = logger.Discard
	}

	db, err := gorm.Open(sqlite.Open("db_books.db"), &gorm.Config{
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
