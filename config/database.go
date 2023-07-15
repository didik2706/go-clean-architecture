package config

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "didik27:Didik.,.88@tcp(localhost:3306)/db_latihan_restful_api_2?parseTime=true")
	if err != nil {
		panic(err)
	}

	// database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}