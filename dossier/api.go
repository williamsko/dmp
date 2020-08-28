package dossier

import (
	"dmp/entity"
	"dmp/usager"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewDossierPayload : for validation of dossier creation API
type NewDossierPayload struct {
	Number string                        `bson:"matricule,omitempty" binding:"required"`
	Usager usager.UsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent  entity.AgentPayloadValidator  `bson:"agent,omitempty" binding:"required"`
}

// NewAntecedentPayload : for validation of antecedent creation API
type NewAntecedentPayload struct {
	AntecedentMedical     string                        `bson:"antecedent_medical,omitempty" binding:"required"`
	AntecedentChirurgical string                        `bson:"antecedent_chirurgical,omitempty" binding:"required"`
	AntecedentFamilial    string                        `bson:"antecedent_familial,omitempty" binding:"required"`
	ModeDeVie             string                        `bson:"mode_de_vie,omitempty" binding:"required"`
	Usager                usager.UsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent                 entity.AgentPayloadValidator  `bson:"agent,omitempty" binding:"required"`
}

//PostDossierAPI : api to create a new empty dmp for usager
func PostDossierAPI(c *gin.Context) {
	var payload NewDossierPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUsager, err := usager.FindUsagerByMatricule(payload.Usager.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	foundAgent, err := entity.FindAgentByMatricule(payload.Agent.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}
	_, err = FindDossierByUsagerID(foundUsager.ID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-already-exists-for-this-usager", "response_code": "100"})
		return
	}
	dossier, err := CreateEmptyDossier(foundUsager, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"response_content": dossier, "response_code": "000"})
}

//PostAntecedentAPI : api to create a new empty dmp for usager
func PostAntecedentAPI(c *gin.Context) {
	var payload NewAntecedentPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUsager, err := usager.FindUsagerByMatricule(payload.Usager.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-usager", "response_code": "100"})
		return
	}
	foundAgent, err := entity.FindAgentByMatricule(payload.Agent.Matricule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-agent", "response_code": "100"})
		return
	}
	dossier, err := FindDossierByUsagerID(foundUsager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-does-not-exist", "response_code": "100"})
		return
	}
	antecedent, err := AddContenuAntecedentUsagerToDossier(dossier, payload, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "antecedent-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"response_content": antecedent, "response_code": "000"})
}
