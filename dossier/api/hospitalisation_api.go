package dossier

import (
	"dmp/dossier"
	repository "dmp/dossier/repository"
	"dmp/entity"
	"dmp/usager"
	"dmp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostHospitalisationAPI : api to add new consultation to dossier
func PostHospitalisationAPI(c *gin.Context) {
	var payload dossier.NewHostpitalisationPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUsager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	agentMatricule, _ := c.Get("agent")
	foundAgent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}
	patientRecord, err := repository.FindPatientRecordByUsagerID(foundUsager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-does-not-exist", "response_code": "100"})
		return
	}
	err = repository.AddContenuHospitalisationToPatientRecord(patientRecord, payload, foundAgent)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "hospitalisation-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, payload)
}

//GetHospitalisationAPI : api to get usager hospitalisation
func GetHospitalisationAPI(c *gin.Context) {
	usager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindPatientRecordByUsagerID(usager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "no-dossier-for-usager", "response_code": "100"})
		return
	}
	// Retreive antecedents usager
	hospitalisationsUsager, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": hospitalisationsUsager, "response_code": "000"})
}
