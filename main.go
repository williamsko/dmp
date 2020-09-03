package main

import (
	"dmp/db"
	api "dmp/dossier/api"
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
		v1.POST("/dossier", api.PostDossierAPI)
		v1.POST("/dossier/antecedent", api.PostAntecedentAPI)
		v1.POST("/dossier/consultation", api.PostConsultationAPI)
		v1.POST("/dossier/hospitalisation", api.PostHospitalisationAPI)
		v1.POST("/dossier/examen", api.PostExamenAPI)

		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	usagers := router.Group("/api/v1/usagers")
	{
		usagers.GET("/", usager.GetAllUsagerAPI)
		usagers.GET("/:matricule", usager.GetUsagerByMatriculeAPI)

	}
	router.Run(":9090")

}
