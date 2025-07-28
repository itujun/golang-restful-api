package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn-golang-restful-api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10) // Set maximum idle connections
	db.SetMaxOpenConns(100) // Set maximum open connections
	db.SetConnMaxLifetime(60 * time.Minute) // Set maximum connection lifetime to 60 minutes
	db.SetConnMaxIdleTime(10 * time.Minute) // Set maximum idle time for connections to 10 minutes

	return db
}