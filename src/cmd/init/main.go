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

	initData(db)
}

func initData(db *gorm.DB) {

	item1 := model.Item{
		Price: 980,
		Name:  "item 1",
	}
	item2 := model.Item{
		Price: 2980,
		Name:  "item 2",
	}
	item3 := model.Item{
		Price: 2980,
		Name:  "item 2",
	}
	items := []*model.Item{
		&item1, &item2, &item3,
	}

	for _, item := range items {
		db.Save(&item)
	}

	cust1 := model.Customer{
		Name:    "name 1",
		Address: "address 1",
	}
	cust2 := model.Customer{
		Name:    "name 2",
		Address: "address 2",
	}
	customers := []*model.Customer{
		&cust1, &cust2,
	}

	for _, cust := range customers {
		db.Save(&cust)
	}

	order1 := model.Order{CustomerID: cust1.ID}
	order2 := model.Order{CustomerID: cust1.ID}
	order3 := model.Order{CustomerID: cust2.ID}
	orders := []*model.Order{
		&order1, &order2, &order3,
	}
	for _, order := range orders {
		db.Save(&order)
	}

	orderItems := []*model.OrderItem{
		{OrderID: order1.ID, ItemID: item1.ID, ItemCount: 5},
		{OrderID: order1.ID, ItemID: item2.ID, ItemCount: 2},
		{OrderID: order2.ID, ItemID: item1.ID, ItemCount: 2},
		{OrderID: order2.ID, ItemID: item3.ID, ItemCount: 1},
		{OrderID: order3.ID, ItemID: item2.ID, ItemCount: 3},
	}

	for _, oi := range orderItems {
		db.Save(&oi)
	}

}
