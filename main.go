package main

import (
	"github.com/arierimbaboemi/bank-auth-service/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	routes.StartServer()
}
