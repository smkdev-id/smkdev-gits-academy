package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UNAMEDB string = "postgres"
	PASSDB  string = "PoSTGreS$$441"
	HOSTDB  string = "localhost"
	DBNAME  string = "todo-db"
)

func Connectdb() *gorm.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOSTDB, UNAMEDB, PASSDB, DBNAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}