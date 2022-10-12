package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Items      []Item `gorm:"foreignKey:ID"`
	CustomerId int
	Customer   Customer
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
