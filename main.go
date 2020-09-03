package main

import (
	"dmp/db"
	dossierApi "dmp/dossier/api"
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
		v1.POST("/dossier", dossierApi.PostDossierAPI)
		v1.POST("/dossier/antecedent", dossierApi.PostAntecedentAPI)
		v1.POST("/dossier/consultation", dossierApi.PostConsultationAPI)
		v1.POST("/dossier/hospitalisation", dossierApi.PostHospitalisationAPI)
		v1.POST("/dossier/examen", dossierApi.PostExamenAPI)

		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	usagers := router.Group("/api/v1/usagers")
	{
		usagers.GET("/", usager.GetAllUsagerAPI)
		usagers.GET("/:matricule", usager.GetUsagerByMatriculeAPI)
		usagers.GET("/:matricule/dossier", dossierApi.GetDossierAPI)
	}

	router.Run(":9090")

}
