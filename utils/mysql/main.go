package mysql

import (
	"database/sql"
	"log"
	"os"
)

var dbPool *sql.DB

// InitConnectionPool is
func InitConnectionPool() {
	const k = "SERVER_ENV"
	v := os.Getenv(k)

	switch v {
	case "dev":
		println("This is development environment nooow!")
		dbPool = initLocalConnectionPool()
	case "stg":
		println("This is staging environment nooow!")
		var err error
		dbPool, err = initSocketConnectionPool()
		if err != nil {
			log.Fatalf("Failed initMySQL() in Staging environment: %v", err)
		}
	default:
		log.Fatalf("Warning: %s environment variable not set.\n", k)
		return
	}

	println("Initilized MySQL!")
}

// DBPool is
func DBPool() *sql.DB {
	return dbPool
}
