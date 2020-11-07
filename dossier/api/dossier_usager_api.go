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
	_, err = repository.FindDossierByUsagerID(foundUsager.ID)
	if err == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-already-exists-for-this-usager")
		return
	}
	numeroDossier, err := repository.CreateEmptyPatientRecord(foundUsager, medecinTraitant, foundAgent)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "dossier-creation-error")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"patient_medical_record_number": numeroDossier	})
}

//GetDossierAPI : api to create a new empty dmp for usager
func GetDossierAPI(c *gin.Context) {
	usager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindDossierByUsagerID(usager.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "no-dossier-for-usager")
		return
	}
	// Retreive antecedents usager
	antecedentsUsager, err := repository.GetAllAntecedentByDossierUsager(&patientRecord)
	consultationsUsager, err := repository.GetAllConsultationsByDossierUsager(&patientRecord)
	hospitalisationsUsager, err := repository.GetAllHospitalisationsByDossierUsager(&patientRecord)
	examensUsager, err := repository.GetAllExamensByDossierUsager(&patientRecord)
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
