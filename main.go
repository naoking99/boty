package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	routes "github.com/naoking99/boty/presentation/router"
	"github.com/naoking99/boty/utils/cloudsql"
	"github.com/naoking99/boty/utils/mysql"
)

func main() {
	const k = "SERVER_ENV"
	v := os.Getenv(k)
	switch v {
	case "dev":
		println("This is development environment nooow!")
		mysql.InitMySQL()
	case "stg":
		println("This is staging environment nooow!")
		cloudsql.InitSocketConnectionPool()
	default:
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}

	routes.HandleRequests()
}
