package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoking99/boty/presentation/router"
	"github.com/naoking99/boty/utils/mysql"
)

func main() {
	mysql.InitConnectionPool()
	router.HandleRequests()
}
