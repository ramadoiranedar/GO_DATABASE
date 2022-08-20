package go_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // minimal create connection db
	db.SetMaxOpenConns(100)                 // maksimal open connection db
	db.SetConnMaxIdleTime(5 * time.Minute)  // maksimal how long traffic idle
	db.SetConnMaxLifetime(60 * time.Minute) // maksimal how long lifetime

	return db
}
