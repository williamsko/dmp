package dossier

import (
	repository "dmp/dossier/repository"
	"dmp/entity"
	"dmp/usager"
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

	foundUsager, err := usager.FindUsagerByMatricule(usagerMatricule)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-usager")
		return
	}
	foundAgent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-agent")
		return
	}

	medecinTraitant, err := entity.FindAgentByMatricule(medecinTraitantMatricule)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-medecin-traitant")
		return
	}
	_, err = repository.FindPatientRecordByUsagerID(foundUsager.ID)
	if err == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-already-exists-for-this-usager")
		return
	}
	numeroDossier, err := repository.CreateEmptyPatientRecord(foundUsager, medecinTraitant, foundAgent)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"patient_medical_record_number": numeroDossier})
}

//GetDossierAPI : get usager dossier and contents by giving usager matricule
func GetDossierAPI(c *gin.Context) {
	usager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindPatientRecordByUsagerID(usager.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "no-dossier-found-for-usager")
		return
	}
	// Retreive antecedents usager
	antecedentsUsager, err := repository.GetAllAntecedentByPatientRecord(&patientRecord)
	consultationsUsager, err := repository.GetAllConsultationsByPatientRecord(&patientRecord)
	hospitalisationsUsager, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)
	examensUsager, err := repository.GetAllExamensByPatientRecord(&patientRecord)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "dossier-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"dossier":          patientRecord,
		"antecedents":      antecedentsUsager,
		"consultations":    consultationsUsager,
		"examens":          examensUsager,
		"hospitalisations": hospitalisationsUsager,
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
	usager, err := usager.FindUsagerByID(patientRecord.Usager)
	agent, err := entity.FindAgentByID(patientRecord.MedecinTraitant)
	antecedents, err := repository.GetAllAntecedentByPatientRecord(&patientRecord)
	consultations, err := repository.GetAllConsultationsByPatientRecord(&patientRecord)
	hospitalisations, err := repository.GetAllHospitalisationsByPatientRecord(&patientRecord)

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"patient_medical_record_number": patientMedicalRecordNumber,
		"usager":                        usager,
		"doctor":                        agent.GetSimpleInformations(),
		"antecedents":                   antecedents,
		"consultations":                 consultations,
		"hospitalisations":              hospitalisations,
	})
}
