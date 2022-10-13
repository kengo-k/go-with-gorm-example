package model

// 100
type Item struct {
	ID    uint
	Price int
	Name  string
}

// 200-
type Customer struct {
	ID      uint
	Name    string
	Address string
}

// 300-
type Order struct {
	ID         uint
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID"`
}

// 400
type OrderItem struct {
	ID        uint
	Order     Order
	OrderID   uint
	Item      Item `gorm:"foreignKey:ItemID"`
	ItemID    uint
	ItemCount int
}
