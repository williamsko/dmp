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

//PostConsultationAPI : api to add new consultation to dossier
func PostConsultationAPI(c *gin.Context) {
	var payload record.NewConsultationPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "unkonwn-patient")
		return
	}
	agentMatricule, _ := c.Get("agent")
	agent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "unkonwn-agent")
		return
	}
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "medical-record-does-not-exist")
		return
	}
	err = repository.AddConsultationToPatientRecord(patientRecord, payload, agent)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "consultation-creation-error")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, payload)
}

//GetConsultationAPI : api to get usager consultation
func GetConsultationAPI(c *gin.Context) {
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "unkonwn-patient")
		return
	}

	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "medical-record-does-not-exist")
		return
	}
	consultationsPatient, err := repository.GetAllConsultationsByPatientRecord(&patientRecord)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "dossier-creation-error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": consultationsPatient, "response_code": "000"})
}
