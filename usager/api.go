package usager

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UsagerPayloadValidator : use to validate usager payload in all models
type UsagerPayloadValidator struct {
	Matricule string `bson:"matricule,omitempty" binding:"required"`
}

//PostUsagerAPI : api to create a new usager
func PostUsagerAPI(c *gin.Context) {
	var usager Usager
	if err := c.BindJSON(&usager); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := FindUsagerByPhoneNumber(usager.PhoneNumber)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "usager-already-exists", "response_code": "100"})
		return
	}
	// We create here a new entry of usager
	newUsager, err := CreateNewUsager(&usager)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": err.Error(), "response_code": "100"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response_content": newUsager, "response_code": "000"})

}
