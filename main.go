package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoking99/boty/infrastracture/config/mysql"
	"github.com/naoking99/boty/presentation/router"
)

func main() {
	mysql.InitConnectionPool()
	router.HandleRequests()
}
