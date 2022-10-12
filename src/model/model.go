package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Items      []Item `gorm:"foreignKey:ID"`
	CustomerID int
	Customer   Customer
}

type OrderItem struct {
	gorm.Model
	OrderID   int
	Order     Order
	ItemID    int
	Item      Item
	ItemCount int
}

type Item struct {
	gorm.Model
	Price int
	Name  string
}

type Customer struct {
	gorm.Model
	Name    string
	Address string
}
