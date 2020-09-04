package dossier

import (
	"dmp/dossier"
	repository "dmp/dossier/repository"
	"dmp/utils"

	"dmp/entity"
	"dmp/usager"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//PostExamenAPI : api to add new consultation to dossier
func PostExamenAPI(c *gin.Context) {
	var payload dossier.NewExamenValidator
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
	dossier, err := repository.FindDossierByUsagerID(foundUsager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-does-not-exist", "response_code": "100"})
		return
	}
	examen, err := repository.AddContenuExamenUsagerToDossier(dossier, payload, foundAgent)
	c.JSON(http.StatusOK, gin.H{"response_content": examen, "response_code": "000"})
}

//GetExamenAPI : api to get usager examen
func GetExamenAPI(c *gin.Context) {
	usager, err := usager.FindUsagerByMatricule(c.Param("matricule"))
	dossierMedical, err := repository.FindDossierByUsagerID(usager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "no-dossier-for-usager", "response_code": "100"})
		return
	}
	examensUsager, err := repository.GetAllExamensByDossierUsager(&dossierMedical)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "dossier-creation-error", "response_code": "100"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": examensUsager, "response_code": "000"})
}

//PatchExamenAPI : updateExamen with results
func PatchExamenAPI(c *gin.Context) {
	var payload dossier.UpdateExamenValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	examen, err := repository.FindUsagerExamenByIdentifiant(c.Param("identifiant"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-examen", "response_code": "100"})
		return
	}
	result, err := repository.UpdateContenuExamen(payload, examen)
	c.JSON(http.StatusOK, gin.H{"response_content": result, "response_code": "000"})
}

//FileUploadAPI : Upload file to usager dossier
func FileUploadAPI(c *gin.Context) {

	idExamen := c.Param("identifiant")
	examen, err := repository.FindUsagerExamenByIdentifiant(idExamen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-examen", "response_code": "100"})
		return
	}
	fmt.Println(examen)

	// single file
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	// Upload the file to gridfs
	err = utils.UploadFile(file, file.Filename)
	fmt.Println(err)

	c.JSON(http.StatusOK, gin.H{"response_content": "OK", "response_code": "000"})
}
