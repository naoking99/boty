package mysql

import (
	"database/sql"
	"log"
	"os"

	"github.com/naoking99/boty/utils/cloudsql"
)

// DBPool is
func DBPool() *sql.DB {
	const k = "SERVER_ENV"
	v := os.Getenv(k)
	switch v {
	case "dev":
		return GetDB()
	case "stg":
		return cloudsql.DBPool()
	default:
		log.Fatalf("Warning: %s environment variable not set.\n", k)
		return nil
	}
}
