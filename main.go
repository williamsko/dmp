package main

import (
	"dmp/db"
	"dmp/dossier"
	"dmp/usager"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDb()
	initRoutes()

}

func initRoutes() {
	router := gin.Default()

	v1 := router.Group("/api/v1/usager")
	{
		v1.POST("/", usager.PostUsagerAPI)
		v1.POST("/dossier", dossier.PostDossierAPI)
		v1.POST("/dossier/antecedent", dossier.PostAntecedentAPI)

		// v1.GET("/", fetchAllTodo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run(":9090")

}
