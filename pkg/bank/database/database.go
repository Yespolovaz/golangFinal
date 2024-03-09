package database

import (
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// MARK: check credentials
	cfg := "host=localhost port=5433 user=postgres dbname=postgres password=123 sslmode=disable"
	db, err := gorm.Open("postgres", cfg)
	helpers.HandleErr(err)
	helpers.HandleErr(err)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)
	DB = db
}
