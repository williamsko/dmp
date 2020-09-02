package dossier

import (
	"dmp/dossier"
	repository "dmp/dossier/repository"

	"dmp/entity"
	"dmp/usager"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PostAntecedentAPI : api to create a new empty dmp for usager
func PostAntecedentAPI(c *gin.Context) {
	var payload dossier.NewAntecedentPayloadValidator
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
	dossierMedical, err := repository.FindDossierByUsagerID(foundUsager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-does-not-exist", "response_code": "100"})
		return
	}
	antecedent, err := repository.AddContenuAntecedentUsagerToDossier(dossierMedical, payload, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "antecedent-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"response_content": antecedent, "response_code": "000"})
}
