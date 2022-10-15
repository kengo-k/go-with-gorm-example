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
		{ID: 201, Number: "0000-0000-0002", UserID: 101},
		{ID: 202, Number: "0000-0000-0003", UserID: 102},
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

func TestHasone(t *testing.T) {
	var users []User
	err := db.
		Model(&User{}).
		Preload("CreditCard").
		Order("id asc").
		Find(&users).
		Error
	if err != nil {
		t.Errorf("err is not nil")
	}

	// test users len
	gotUsersLen := len(users)
	wantUsersLen := 3
	if gotUsersLen != wantUsersLen {
		t.Errorf("users len: got=%v, want=%v", gotUsersLen, wantUsersLen)
	}

	// test cards len
	var cards []CreditCard
	err = db.Model(&CreditCard{}).Find(&cards).Error
	if err != nil {
		t.Errorf("err is not nil")
	}
	gotCardsLen := len(cards)
	wantCardsLen := 3
	if gotCardsLen != wantCardsLen {
		t.Errorf("cards len: got=%v, want=%v", gotCardsLen, wantCardsLen)
	}

	// test users and cards relations
	user1 := users[0]
	user1Card := user1.CreditCard
	user2 := users[1]
	user2Card := user2.CreditCard
	user3 := users[2]
	user3Card := user3.CreditCard

	checkCard := func(got string, actual string) {
		if got != actual {
			t.Errorf("card: got=%v, actual=%v", got, actual)
		}
	}
	checkCard(user1Card.Number, "0000-0000-0001")
	checkCard(user2Card.Number, "0000-0000-0002")
	checkCard(user3Card.Number, "0000-0000-0003")
}
