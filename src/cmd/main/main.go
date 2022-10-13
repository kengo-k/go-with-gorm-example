package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sample-go-with-gorm.com/src/service"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()

	service := &service.Service{}
	/*orders := */ service.SelectAllOrders(db)
	//service.SelectAllOrderItems(db)

	// fmt.Printf("order count: %d\n", len(orders))
	// for _, order := range orders {
	// 	fmt.Printf("--------------\n")
	// 	fmt.Printf("  order id: %d\n", order.ID)
	// 	orderItems := order.OrderItems
	// 	fmt.Printf("    order items count: %d\n", len(orderItems))
	// 	for _, oi := range orderItems {
	// 		fmt.Printf("      item id: %d\n", oi.Item.ID)
	// 	}
	// 	cust := order.Customer
	// 	fmt.Printf("    customer id: %d\n", cust.ID)
	// }
}
