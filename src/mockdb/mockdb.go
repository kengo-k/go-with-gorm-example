package mockdb

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, sqlmock.Sqlmock) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		panic("failed to create sqlmock.New()")
	}
	dialect := postgres.New(postgres.Config{
		Conn: conn,
	})
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic("failed to create gorm.Open()")
	}
	return db, mock
}
