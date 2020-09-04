package dossier

import (
	"dmp/dossier"
	repository "dmp/dossier/repository"

	"dmp/entity"
	"dmp/usager"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PostDossierAPI : api to create a new empty dmp for usager
func PostDossierAPI(c *gin.Context) {
	var payload dossier.NewDossierPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUsager, err := usager.FindUsagerByMatricule(payload.Usager.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	foundAgent, err := entity.FindAgentByMatricule(payload.Agent.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}
	_, err = repository.FindDossierByUsagerID(foundUsager.ID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-already-exists-for-this-usager", "response_code": "100"})
		return
	}
	dossierMedical, err := repository.CreateEmptyDossier(foundUsager, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": dossierMedical, "response_code": "000"})
}

//GetDossierAPI : api to create a new empty dmp for usager
func GetDossierAPI(c *gin.Context) {
	usager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	dossierMedical, err := repository.FindDossierByUsagerID(usager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "no-dossier-for-usager", "response_code": "100"})
		return
	}
	// Retreive antecedents usager
	antecedentsUsager, err := repository.GetAllAntecedentByDossierUsager(&dossierMedical)
	consultationsUsager, err := repository.GetAllConsultationsByDossierUsager(&dossierMedical)
	hospitalisationsUsager, err := repository.GetAllHospitalisationsByDossierUsager(&dossierMedical)
	examensUsager, err := repository.GetAllExamensByDossierUsager(&dossierMedical)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": gin.H{
		"dossier":          dossierMedical,
		"antecedents":      antecedentsUsager,
		"consultations":    consultationsUsager,
		"examens":          examensUsager,
		"hospitalisations": hospitalisationsUsager,
	},
		"response_code": "000",
	})
}
