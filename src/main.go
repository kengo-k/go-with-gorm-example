package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	newProduct := Product{
		Code:  "0001",
		Price: 980,
	}
	db.Create(&newProduct)

	// Read
	var product Product
	db.First(&product, newProduct.ID)
	db.First(&product, "code = ?", "0001")
	fmt.Printf("created product: %#v\n", newProduct)
	fmt.Printf("selected product: %#v\n", product)
}
