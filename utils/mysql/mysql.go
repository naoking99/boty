package mysql

import (
	"database/sql"
	"time"
)

func initLocalMySQL() *sql.DB {
	dbPool, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/boty-dev")
	if err != nil {
		panic(err.Error())
	}

	dbPool.SetConnMaxLifetime(time.Minute * 3)
	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxIdleConns(10)

	return dbPool
}
