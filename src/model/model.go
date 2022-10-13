package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	CustomerID int
	Customer   Customer `gorm:"foreignKey:ID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   int
	Item      Item `gorm:"foreignKey:ID"`
	ItemID    int
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
