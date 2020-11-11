package routes

import (
	agentApi "dmp/entity/api"
	"dmp/patient"
	dossierApi "dmp/record/api"
	"dmp/utils"
	"log"
	"net/http"

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
		c.Set("agent", matricule)
		c.Next()
	}
}

// CorsMiddleware : cors middleware
func CorsMiddleware(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

// SetupRoutes : setup routes for project
func SetupRoutes() *gin.Engine {

	router := gin.Default()
	// router.Use(CorsMiddleware)

	dossierRouter := router.Group("/api/v1/record")
	dossierRouter.Use(TokenAuthMiddleware())
	{
		dossierRouter.GET("/", dossierApi.SearchDossierAPI)

	}

	usagerRouter := router.Group("/api/v1/patient")
	usagerRouter.Use(TokenAuthMiddleware())

	{
		usagerRouter.POST("/", patient.PostUsagerAPI)
		usagerRouter.POST("/:matricule/record", dossierApi.PostDossierAPI)

		usagerRouter.PUT("/:matricule/record/antecedent", dossierApi.PostAntecedentAPI)
		usagerRouter.PUT("/:matricule/record/consultation", dossierApi.PostConsultationAPI)
		usagerRouter.PUT("/:matricule/record/hospitalisation", dossierApi.PostHospitalisationAPI)
		usagerRouter.PUT("/:matricule/record/examen", dossierApi.PostExamenAPI)

		usagerRouter.GET("/:matricule", patient.GetPatientByMatriculeAPI)
		usagerRouter.GET("/:matricule/record", dossierApi.GetDossierAPI)
		usagerRouter.GET("/:matricule/record/antecedents", dossierApi.GetAntecedentsAPI)
		usagerRouter.GET("/:matricule/record/consultations", dossierApi.GetConsultationAPI)
		usagerRouter.GET("/:matricule/record/hospitalisations", dossierApi.GetHospitalisationAPI)
		usagerRouter.GET("/:matricule/record/examens", dossierApi.GetExamenAPI)

		usagerRouter.PATCH("/:matricule/record/examen/:identifiant", dossierApi.PatchExamenAPI)
	}
	usagersRoute := router.Group("/api/v1/patients")
	usagersRoute.Use(TokenAuthMiddleware())

	{
		usagersRoute.GET("/", patient.GetPatientsAPI)

	}

	fileUploadRouter := router.Group("/api/v1/upload")
	fileUploadRouter.Use(TokenAuthMiddleware())
	{
		fileUploadRouter.POST("/patient/:matricule/record/examen/:identifiant", dossierApi.FileUploadAPI)

	}
	fileDownloadRouter := router.Group("/api/v1/download")
	fileDownloadRouter.Use(TokenAuthMiddleware())
	{
		fileDownloadRouter.GET("/patient/:matricule/record/examen/:identifiant/file", dossierApi.FileDownloadAPI)

	}

	agentRouter := router.Group("/api/v1/agent")
	{
		agentRouter.POST("/signin", agentApi.AgentLoginAPI)
	}

	log.Print("Starting DMP Server ...")

	return router

}
