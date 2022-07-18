package repository

import (
	"go-test-web-app/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect a database/n")
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
