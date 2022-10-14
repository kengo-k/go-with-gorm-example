package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(dbName string) *gorm.DB {

	// delete database file if already exists
	_, statErr := os.Stat(dbName)
	if statErr == nil {
		removeErr := os.Remove(dbName)
		if removeErr != nil {
			panic("failed to delete database")
		}
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to create database")
	}
	return db
}
