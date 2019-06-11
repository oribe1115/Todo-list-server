package model

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	db          *gorm.DB
	detabaseURL string
)

// DBとの接続
func EstablishConnection() (*gorm.DB, error) {
	detabaseURL = os.Getenv("DATABASE_URL")
	_db, err := gorm.Open("postgres", detabaseURL)
	if err != nil {
		return nil, errors.New("faild to connect to DB")
	}
	db = _db

	return db, nil
}
