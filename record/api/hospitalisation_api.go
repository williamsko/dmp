package record

import (
	"dmp/record"
	repository "dmp/record/repository"
	"dmp/entity"
	"dmp/patient"
	"dmp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostHospitalisationAPI : api to add new consultation to dossier
func PostHospitalisationAPI(c *gin.Context) {
	var payload record.NewHostpitalisationPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-patient")
		return
	}
	agentMatricule, _ := c.Get("agent")
	agent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-agent")
		return
	}
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "medical-record-does-not-exist")
		return
	}
	err = repository.AddContenuHospitalisationToPatientRecord(patientRecord, payload, agent)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "hospitalisation-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, payload)
}

//GetHospitalisationAPI : api to get usager hospitalisation
func GetHospitalisationAPI(c *gin.Context) {
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "medical-record-does-not-exist")
		return
	}
	// Retreive antecedents usager
	hospitalisationsPatient, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)

	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "dossier-creation-error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": hospitalisationsPatient, "response_code": "000"})
}
