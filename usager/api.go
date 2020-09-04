package usager

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewUsagerPayloadValidator : use to validate new usager payload in all models
type NewUsagerPayloadValidator struct {
	FirstName             string `json:"first_name,omitempty" binding:"required"`
	LastName              string `json:"last_name,omitempty" binding:"required"`
	Address               string `json:"address,omitempty" binding:"required"`
	PhoneNumber           string `json:"phone_number,omitempty" binding:"required"`
	IdentityNumber        string `json:"identity_number,omitempty" binding:"required"`
	TypeDocument          string `json:"type_document,omitempty" binding:"required"`
	Sexe                  string `json:"sexe,omitempty" binding:"required"`
	SituationMatrimoniale string `json:"situation_matrimoniale,omitempty" binding:"required"`
}

// FindUsagerPayloadValidator : use to find usager payload
type FindUsagerPayloadValidator struct {
	Matricule string `json:"matricule,omitempty" unique:"true" binding:"required"`
}

//PostUsagerAPI : api to create a new usager
func PostUsagerAPI(c *gin.Context) {
	var payload NewUsagerPayloadValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := FindUsagerByPhoneNumber(payload.PhoneNumber)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "usager-already-exists", "response_code": "100"})
		return
	}
	newUsager, err := CreateNewUsager(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": err.Error(), "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": newUsager, "response_code": "000"})
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
