package patient

import (
	"dmp/entity"
	"dmp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUsagerPayloadValidator : use to validate new usager payload in all models
type NewUsagerPayloadValidator struct {
	FirstName                  string                     `json:"first_name" binding:"required"`
	LastName                   string                     `json:"last_name" binding:"required"`
	Address                    string                     `json:"address" binding:"required"`
	PhoneNumber                string                     `json:"phone_number" binding:"required"`
	IdentityNumber             string                     `json:"identity_number" binding:"required"`
	TypeDocument               string                     `json:"type_document" binding:"required"`
	Sexe                       string                     `json:"sexe" binding:"required"`
	SituationMatrimoniale      string                     `json:"situation_matrimoniale" binding:"required"`
	PersonneaPrevenirValidator PersonneaPrevenirValidator `json:"personne_a_prevenir" binding:"required"`
}

// PersonneaPrevenirValidator : use to validate personne a prevenir payload
type PersonneaPrevenirValidator struct {
	FirstName          string `json:"first_name" binding:"required"`
	LastName           string `json:"last_name" binding:"required"`
	Address            string `json:"address" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	Sexe               string `json:"sexe" binding:"required"`
	RelationWithUsager string `json:"relation_with_usager" binding:"required"`
}

// FindUsagerPayloadValidator : use to find usager payload
type FindUsagerPayloadValidator struct {
	Matricule string `json:"matricule,omitempty" unique:"true" binding:"required"`
}

//PostUsagerAPI : API creation of new usager
func PostUsagerAPI(c *gin.Context) {
	var payload Patient
	agentMatricule, _ := c.Get("agent")
	if err := c.BindJSON(&payload); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-agent")
		return
	}

	_, err = FindPatientByPhoneNumber(payload.PhoneNumber)
	if err == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "usager-already-exists")
		return
	}
	patient, err := CreateNewUsager(&payload)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"usager": patient})
}

//GetPatientsAPI : api get all usagers
func GetPatientsAPI(c *gin.Context) {
	usagers, err := GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": err.Error(), "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": usagers, "response_code": "000"})
}

//GetPatientByMatriculeAPI : api get  usager
func GetPatientByMatriculeAPI(c *gin.Context) {
	patient, err := FindPatientByMatricule(c.Param("matricule"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-patient")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, patient)
}
