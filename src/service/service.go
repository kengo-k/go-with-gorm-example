package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sample-go-with-gorm.com/model"
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
		Preload(clause.Associations).
		Find(&orders)

	return orders
}

func (s *Service) SelectAllOrderItems(db *gorm.DB) []model.OrderItem {
	orderItems := []model.OrderItem{}
	db.
		Preload(clause.Associations).
		Find(&orderItems)

	return orderItems
}
