package mysql

import (
	"database/sql"
	"time"
)

var db *sql.DB

// InitMySQL is
func InitMySQL() {
	_db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/boty-dev")
	if err != nil {
		panic(err.Error())
	}

	db = _db

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	println("Initilized MySQL!")
}

// GetDB is
func GetDB() *sql.DB {
	return db
}
