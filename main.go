package main

import (
	"dmp/db"
	"dmp/routes"
)

func main() {

	db.ConnectDb()
	routes.SetupRoutes()

}
