package routes

import (
	dossierApi "dmp/dossier/api"
	"dmp/usager"
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRoutes : setup routes for project
func SetupRoutes() *gin.Engine {

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
		usagerRouter.GET("/:matricule/dossier/hospitalisations", dossierApi.GetHospitalisationAPI)
		usagerRouter.GET("/:matricule/dossier/examens", dossierApi.GetExamenAPI)

		usagerRouter.PATCH("/:matricule/dossier/examen/:identifiant", dossierApi.PatchExamenAPI)
	}
	usagersRoute := router.Group("/api/v1/usagers")
	{
		usagersRoute.GET("/", usager.GetAllUsagerAPI)

	}

	fileUploadRouter := router.Group("/api/v1/upload")
	{
		fileUploadRouter.POST("/usager/:matricule/dossier/examen/:identifiant", dossierApi.FileUploadAPI)

	}
	fileDownloadRouter := router.Group("/api/v1/download")
	{
		fileDownloadRouter.GET("/usager/:matricule/dossier/examen/:identifiant/file", dossierApi.FileDownloadAPI)

	}

	log.Print("Starting DMP Server ...")

	return router

}
