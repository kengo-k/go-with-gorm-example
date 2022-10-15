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

func TestMain(m *testing.M) {
	db = database.New("hasone.db")
	db.AutoMigrate(&User{}, &CreditCard{})
	m.Run()
}

func TestIt(t *testing.T) {
}
