package main

import (
	"dmp/db"
	"dmp/routes"
)

func main() {

	

	db.ConnectDb()

	router := routes.SetupRoutes()

	router.Run(":9090")

}
