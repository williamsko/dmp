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

		v1.PUT("/dossier/antecedent", dossierApi.PostAntecedentAPI)
		v1.PUT("/dossier/consultation", dossierApi.PostConsultationAPI)
		v1.PUT("/dossier/hospitalisation", dossierApi.PostHospitalisationAPI)
		v1.PUT("/dossier/examen", dossierApi.PostExamenAPI)

		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	usagers := router.Group("/api/v1/usagers")
	{
		usagers.GET("/", usager.GetAllUsagerAPI)
		usagers.GET("/:matricule", usager.GetUsagerByMatriculeAPI)
		usagers.GET("/:matricule/dossier", dossierApi.GetDossierAPI)

		usagers.GET("/:matricule/dossier/antecedents", dossierApi.GetAntecedentAPI)
		usagers.GET("/:matricule/dossier/consultations", dossierApi.GetConsultationAPI)
		usagers.GET("/:matricule/dossier/hospitalisations", dossierApi.GetHispitalisationAPI)
		usagers.GET("/:matricule/dossier/examens", dossierApi.GetExamenAPI)

		usagers.PATCH("/:matricule/dossier/examens/:identifiant", dossierApi.PatchExamenAPI)

	}

	router.Run(":9090")

}
