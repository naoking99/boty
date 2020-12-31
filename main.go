package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoking99/boty/presentation/routes"
	"github.com/naoking99/boty/utils/mysql"
)

func main() {
	mysql.InitMySQL()
	routes.HandleRequests()
}
