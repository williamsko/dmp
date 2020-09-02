package dossier

import (
	"dmp/dossier"
	repository "dmp/dossier/repository"

	"dmp/entity"
	"dmp/usager"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PostExamenAPI : api to add new consultation to dossier
func PostExamenAPI(c *gin.Context) {
	var payload dossier.NewExamenValidator
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
	dossier, err := repository.FindDossierByUsagerID(foundUsager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-does-not-exist", "response_code": "100"})
		return
	}

	examen, err := repository.AddContenuExamenUsagerToDossier(dossier, payload, foundAgent)

	c.JSON(http.StatusBadRequest, gin.H{"response_content": examen, "response_code": "000"})
}
