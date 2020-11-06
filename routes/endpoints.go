package routes

import (
	dossierApi "dmp/dossier/api"
	agentApi "dmp/entity/api"
	"dmp/usager"
	"dmp/utils"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware : token authentication middleware
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		matricule := utils.ExtractTokenAuth(c.Request)
		if matricule == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "api-token-required-or-expired")
			return
		}
		log.Print(matricule)
		c.Set("agent",matricule)
		c.Next()
	}
}

// SetupRoutes : setup routes for project
func SetupRoutes() *gin.Engine {

	router := gin.Default()
	usagerRouter := router.Group("/api/v1/usager")
	usagerRouter.Use(TokenAuthMiddleware())

	{
		usagerRouter.POST("/", usager.PostUsagerAPI)
		usagerRouter.POST("/:matricule/dossier", dossierApi.PostDossierAPI)

		usagerRouter.PUT("/:matricule/dossier/antecedent", dossierApi.PostAntecedentAPI)
		usagerRouter.PUT("/:matricule/dossier/consultation", dossierApi.PostConsultationAPI)
		usagerRouter.PUT("/:matricule/dossier/hospitalisation", dossierApi.PostHospitalisationAPI)
		usagerRouter.PUT("/:matricule/dossier/examen", dossierApi.PostExamenAPI)

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

	agentRouter := router.Group("/api/v1/agent")
	{
		agentRouter.POST("/signin", agentApi.AgentLoginAPI)

	}

	log.Print("Starting DMP Server ...")

	return router

}
