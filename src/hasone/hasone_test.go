package hasone

import (
	"testing"

	"gorm.io/gorm"
	database "sample-go-with-gorm.com/db"
)

type User struct {
	ID         uint
	CreditCard CreditCard
}

type CreditCard struct {
	ID     uint
	Number string
	UserID uint
}

var db *gorm.DB

func initDB() {
	db = database.New("hasone.db")
	db.AutoMigrate(&User{}, &CreditCard{})
	users := []*User{
		{ID: 100}, {ID: 101}, {ID: 102},
	}
	cards := []*CreditCard{
		{ID: 200, Number: "0000-0000-0001", UserID: 100},
		{ID: 201, Number: "0000-0000-0002", UserID: 100},
		{ID: 202, Number: "0000-0000-0003", UserID: 100},
		{ID: 203, Number: "0000-0000-0004", UserID: 101},
		{ID: 204, Number: "0000-0000-0005", UserID: 102},
		{ID: 205, Number: "0000-0000-0006", UserID: 102},
	}

	for _, u := range users {
		db.Save(u)
	}

	for _, c := range cards {
		db.Save(c)
	}
}

func TestMain(m *testing.M) {
	initDB()
	m.Run()
}

func TestIt(t *testing.T) {
}
