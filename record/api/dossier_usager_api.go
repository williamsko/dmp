package record

import (
	repository "dmp/record/repository"
	"dmp/entity"
	"dmp/patient"
	"dmp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type affectation struct {
	MedecinTraitantMatricule string `json:"affectation" binding:"required"`
}

//PostDossierAPI : api to create a new empty dmp for usager
func PostDossierAPI(c *gin.Context) {
	agentMatricule, _ := c.Get("agent")
	usagerMatricule := c.Param("matricule")

	var payload affectation
	if err := c.BindJSON(&payload); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	medecinTraitantMatricule := payload.MedecinTraitantMatricule

	patient, err := patient.FindPatientByMatricule(usagerMatricule)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-patient")
		return
	}
	agent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-agent")
		return
	}

	medecinTraitant, err := entity.FindAgentByMatricule(medecinTraitantMatricule)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-medecin-traitant")
		return
	}
	_, err = repository.FindRecordByPatientD(patient.ID)
	if err == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-already-exists-for-this-usager")
		return
	}
	numeroDossier, err := repository.CreateEmptyPatientRecord(patient, medecinTraitant, agent)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"patient_medical_record_number": numeroDossier})
}

//GetDossierAPI : get usager dossier and contents by giving usager matricule
func GetDossierAPI(c *gin.Context) {
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "no-dossier-found-for-usager")
		return
	}
	agent, err := entity.FindAgentByID(patientRecord.MedecinTraitant)
	antecedents, err := repository.GetAntecedentsByPatientRecord(&patientRecord)
	consultations, err := repository.GetAllConsultationsByPatientRecord(&patientRecord)
	hospitalisations, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)
	examens, err := repository.GetAllExamensByPatientRecord(&patientRecord)

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"patient_medical_record_number": patientRecord.Number,
		"usager":                        patient,
		"doctor":                        agent.GetSimpleInformations(),
		"antecedents":                   antecedents,
		"consultations":                 consultations,
		"hospitalisations":              hospitalisations,
		"examens":                       examens,
	})
}

//SearchDossierAPI : api to create a new empty dmp for usager
func SearchDossierAPI(c *gin.Context) {
	patientMedicalRecordNumber := c.Query("patient_medical_record_number")
	if len(patientMedicalRecordNumber) <= 0 {
		utils.RespondWithError(c, http.StatusNotFound, "empty-query")
		return
	}
	patientRecord, err := repository.FindPatientRecordByNumber(patientMedicalRecordNumber)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "no-dossier-found-for-usager")
		return
	}
	patient, err := patient.FindPatientByID(patientRecord.Usager)
	agent, err := entity.FindAgentByID(patientRecord.MedecinTraitant)
	antecedents, err := repository.GetAntecedentsByPatientRecord(&patientRecord)
	consultations, err := repository.GetAllConsultationsByPatientRecord(&patientRecord)
	hospitalisations, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)
	examens, err := repository.GetAllExamensByPatientRecord(&patientRecord)

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"patient_medical_record_number": patientMedicalRecordNumber,
		"usager":                        patient,
		"doctor":                        agent.GetSimpleInformations(),
		"antecedents":                   antecedents,
		"consultations":                 consultations,
		"hospitalisations":              hospitalisations,
		"examens":                       examens,
	})
}
