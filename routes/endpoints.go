package routes

import (
	dossierApi "dmp/dossier/api"
	"dmp/usager"
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRoutes : setup routes for project
func SetupRoutes() {

	router := gin.Default()
	usagerDossierCreationRouter := router.Group("/api/v1/usager")
	{
		usagerDossierCreationRouter.POST("/", usager.PostUsagerAPI)
		usagerDossierCreationRouter.POST("/dossier", dossierApi.PostDossierAPI)

		usagerDossierCreationRouter.PUT("/dossier/antecedent", dossierApi.PostAntecedentAPI)
		usagerDossierCreationRouter.PUT("/dossier/consultation", dossierApi.PostConsultationAPI)
		usagerDossierCreationRouter.PUT("/dossier/hospitalisation", dossierApi.PostHospitalisationAPI)
		usagerDossierCreationRouter.PUT("/dossier/examen", dossierApi.PostExamenAPI)

		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	usagerDossierConsultationRouter := router.Group("/api/v1/usagers")
	{
		usagerDossierConsultationRouter.GET("/", usager.GetAllUsagerAPI)
		usagerDossierConsultationRouter.GET("/:matricule", usager.GetUsagerByMatriculeAPI)
		usagerDossierConsultationRouter.GET("/:matricule/dossier", dossierApi.GetDossierAPI)

		usagerDossierConsultationRouter.GET("/:matricule/dossier/antecedents", dossierApi.GetAntecedentAPI)
		usagerDossierConsultationRouter.GET("/:matricule/dossier/consultations", dossierApi.GetConsultationAPI)
		usagerDossierConsultationRouter.GET("/:matricule/dossier/hospitalisations", dossierApi.GetHispitalisationAPI)
		usagerDossierConsultationRouter.GET("/:matricule/dossier/examens", dossierApi.GetExamenAPI)

		usagerDossierConsultationRouter.PATCH("/:matricule/dossier/examens/:identifiant", dossierApi.PatchExamenAPI)
	}

	fileUploadRouter := router.Group("/api/v1/file")
	{
		fileUploadRouter.POST("/usager/:matricule/dossier/examens/:identifiant", dossierApi.FileUploadAPI)

	}
	log.Print("This is our first log message in Go.")

	router.Run(":9090")

}
