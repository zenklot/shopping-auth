package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/shopping")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 10)
	return db
}
