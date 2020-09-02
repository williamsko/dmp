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
	c.JSON(http.StatusBadRequest, gin.H{"response_content": dossierMedical, "response_code": "000"})
}
