package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db  *sqlx.DB
	err error
)

func Init() {
	db, err = sqlx.Open("postgres", "user=gotest password=gotest!password dbname=gotest sslmode=disable")
	if err != nil {
		panic("Failed to connect a database/n")
	}
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
