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
	db.AutoMigrate(&model.OrderItem{}, &model.Order{}, &model.Customer{}, &model.Item{})

	items := []*model.Item{
		{
			Price: 980,
			Name:  "item 1",
		},
		{
			Price: 2980,
			Name:  "item 2",
		},
		{
			Price: 98,
			Name:  "item 3",
		},
	}

	for _, item := range items {
		db.Save(item)
	}

}
