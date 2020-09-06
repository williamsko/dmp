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
	usagerRouter := router.Group("/api/v1/usager")
	{
		usagerRouter.POST("/", usager.PostUsagerAPI)
		usagerRouter.POST("/dossier", dossierApi.PostDossierAPI)

		usagerRouter.PUT("/dossier/antecedent", dossierApi.PostAntecedentAPI)
		usagerRouter.PUT("/dossier/consultation", dossierApi.PostConsultationAPI)
		usagerRouter.PUT("/dossier/hospitalisation", dossierApi.PostHospitalisationAPI)
		usagerRouter.PUT("/dossier/examen", dossierApi.PostExamenAPI)

		usagerRouter.GET("/:matricule", usager.GetUsagerByMatriculeAPI)
		usagerRouter.GET("/:matricule/dossier", dossierApi.GetDossierAPI)
		usagerRouter.GET("/:matricule/dossier/antecedents", dossierApi.GetAntecedentAPI)
		usagerRouter.GET("/:matricule/dossier/consultations", dossierApi.GetConsultationAPI)
		usagerRouter.GET("/:matricule/dossier/hospitalisations", dossierApi.GetHispitalisationAPI)
		usagerRouter.GET("/:matricule/dossier/examens", dossierApi.GetExamenAPI)

		usagerRouter.PATCH("/:matricule/dossier/examens/:identifiant", dossierApi.PatchExamenAPI)
	}
	usagersRoute := router.Group("/api/v1/usagers")
	{
		usagersRoute.GET("/", usager.GetAllUsagerAPI)

	}

	fileUploadRouter := router.Group("/api/v1/file")
	{
		fileUploadRouter.POST("/usager/:matricule/dossier/examens/:identifiant", dossierApi.FileUploadAPI)

	}
	log.Print("This is our first log message in Go.")

	router.Run(":9090")

}
