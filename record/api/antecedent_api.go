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

//PostAntecedentAPI : api to create a new empty dmp for usager
func PostAntecedentAPI(c *gin.Context) {
	var payload record.NewAntecedentPayloadValidator
	agentMatricule, _ := c.Get("agent")

	if err := c.BindJSON(&payload); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "unkonwn-patient")
		return
	}
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
	err = repository.AddAntecedentToPatientRecord(patientRecord, payload, agent)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "adding-antecedent-to-patient-record-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, payload)
}

//GetAntecedentsAPI : api to get usager antecedent
func GetAntecedentsAPI(c *gin.Context) {
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "medical-record-does-not-exist")
		return
	}
	antecedents, err := repository.GetAntecedentsByPatientRecord(&patientRecord)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "dossier-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, antecedents)
}
