package dossier

import (
	repository "dmp/dossier/repository"
	"dmp/entity"
	"dmp/usager"
	"net/http"

	"github.com/gin-gonic/gin"
)

type affectation struct {
	MedecinTraitantMatricule  string  `json:"affectation" binding:"required"`
}

//PostDossierAPI : api to create a new empty dmp for usager
func PostDossierAPI(c *gin.Context) {
	agentMatricule, _ := c.Get("agent")
	usagerMatricule := c.Param("matricule")

	var payload affectation
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medecinTraitantMatricule := payload.MedecinTraitantMatricule
	
	foundUsager, err := usager.FindUsagerByMatricule(usagerMatricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	foundAgent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}

	medecinTraitant, err := entity.FindAgentByMatricule(medecinTraitantMatricule)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-medecin-traitant", "response_code": "100"})
		return
	}
	_, err = repository.FindDossierByUsagerID(foundUsager.ID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-already-exists-for-this-usager", "response_code": "100"})
		return
	}
	numeroDossier, err := repository.CreateEmptyDossier(foundUsager, medecinTraitant, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"numero_dossier": numeroDossier, "response_code": "000"})
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
