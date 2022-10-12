package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sample-go-with-gorm.com/src/model"
)

// initial database, migrate and insert test data
func main() {
	dbName := "test.db"

	// delete database file if already exists
	_, statErr := os.Stat(dbName)
	if statErr == nil {
		removeErr := os.Remove(dbName)
		if removeErr != nil {
			panic("failed to delete database")
		}
	}

	// create new database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to create database")
	}

	// execute migrate
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.OrderItem{})

}
