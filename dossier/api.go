package dossier

import (
	"dmp/entity"
	"dmp/usager"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PostDossierAPI : api to create a new empty dmp for usager
func PostDossierAPI(c *gin.Context) {
	var payload NewDossierPayloadValidator
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
	var payload NewAntecedentPayloadValidator
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

//PostConsultationAPI : api to add new consultation to dossier
func PostConsultationAPI(c *gin.Context) {
	var payload NewConsultationPayloadValidator
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
	consultation, err := AddContenuConsultationUsagerToDossier(dossier, payload, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "antecedent-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"response_content": consultation, "response_code": "000"})
}

//PostHospitalisationAPI : api to add new consultation to dossier
func PostHospitalisationAPI(c *gin.Context) {
	var payload NewHostpitalisationPayloadValidator
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
	consultation, err := AddContenuHospitalisationUsagerToDossier(dossier, payload, foundAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "antecedent-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"response_content": consultation, "response_code": "000"})
}

//PostExamenAPI : api to add new consultation to dossier
func PostExamenAPI(c *gin.Context) {
	var payload NewExamenValidator
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

	fmt.Println(foundAgent)
	fmt.Println(foundUsager)
	fmt.Println(dossier)

	c.JSON(http.StatusBadRequest, gin.H{"response_content": "OK", "response_code": "000"})
}
