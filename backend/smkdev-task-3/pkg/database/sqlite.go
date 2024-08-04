package database

import (
	"fmt"

	"github.com/FuadGrimaldi/bookstore-api/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connection to SQLite database
func ConnectToSqlite(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Sqlite
	fmt.Printf("Using database path: %s\n", dsn) // Log the path
	var log logger.Interface
	if cfg.Env == "dev" {
		log = logger.Default.LogMode(logger.Info)
		} else {
			log = logger.Default.LogMode(logger.Silent)
		}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the SQLite database: %w", err)
	}
	// Check if the connection is valid
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the SQLite database: %w", err)
	}
	return db, nil
}