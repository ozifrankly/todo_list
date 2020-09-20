package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

//Connect init database
func Connect() *gorm.DB {

	username := os.Getenv("USERNAME")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable ", dbHost, username, dbName)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
	return db
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
