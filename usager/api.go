package usager

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

//PostUsagerAPI : api to create a new usager
func PostUsagerAPI(c *gin.Context) {
	var payload Usager
	agentMatricule, _ := c.Get("agent")
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}

	_, err = FindUsagerByPhoneNumber(payload.PhoneNumber)
	if err == nil {
		utils.RespondWithError(c, http.StatusBadRequest, "usager-already-exists")
		return
	}
	createdUsager, err := CreateNewUsager(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": err.Error(), "response_code": "100"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usager": createdUsager, "response_code": "000"})
}

//GetAllUsagerAPI : api get all usagers
func GetAllUsagerAPI(c *gin.Context) {
	usagers, err := GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": err.Error(), "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": usagers, "response_code": "000"})
}

//GetUsagerByMatriculeAPI : api get  usager
func GetUsagerByMatriculeAPI(c *gin.Context) {
	usager, err := FindUsagerByMatricule(c.Param("matricule"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": usager, "response_code": "000"})
}
