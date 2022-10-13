package service

import (
	"gorm.io/gorm"
	"sample-go-with-gorm.com/src/model"
)

type Service struct{}

func (s *Service) SelectAllItems(db *gorm.DB) []model.Item {
	items := []model.Item{}
	db.Find(&items)
	return items
}

func (s *Service) SelectAllOrders(db *gorm.DB) []model.Order {
	orders := []model.Order{}
	db.
		Preload("Customer").
		Preload("OrderItems").
		Preload("Item").
		Find(&orders)
	return orders
}
