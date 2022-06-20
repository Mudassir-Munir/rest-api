package db

import (
	"fmt"

	"github.com/Mudassir-Munir/rest-api/pkg/modles"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres-container"
	password = "postgres"
	dbname   = "todoList-db"
)

//db connection function
func Init() *gorm.DB {
	// connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&modles.Book{})
	// defer db.Close()

	return db

}
