package main

import (
	_ "github.com/go-sql-driver/mysql"
	routes "github.com/naoking99/boty/presentation/router"
	"github.com/naoking99/boty/utils/cloudsql"
)

func main() {
	// mysql.InitMySQL()
	cloudsql.InitSocketConnectionPool()
	// cloudsql.InitTCPConnectionPool()
	routes.HandleRequests()
}
