package app

import (
	"database/sql"
	"happy-product-service/helper"
	"time"
)

func NewDB(dbDriver, dbSource string) *sql.DB {
	db, err := sql.Open(dbDriver, dbSource)
	helper.DoPanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
