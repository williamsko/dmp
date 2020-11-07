package dossier

import (
	"bytes"
	"dmp/dossier"
	repository "dmp/dossier/repository"
	"dmp/entity"
	"dmp/usager"
	"dmp/utils"
	"github.com/gin-gonic/gin"
	"image"
	jpeg "image/jpeg"
	"io"
	"log"
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
	dossier, err := repository.FindPatientRecordByUsagerID(foundUsager.ID)
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
	patientRecord, err := repository.FindPatientRecordByUsagerID(usager.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "no-dossier-for-usager", "response_code": "100"})
		return
	}
	examensUsager, err := repository.GetAllExamensByDossierUsager(&patientRecord)
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
	log.Println(examen)

	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to gridfs
	fileID, err := utils.UploadFile(file, file.Filename)
	_, err = repository.UpdateContenuExamenWithFile(examen, fileID)
	log.Println(err)

	c.JSON(http.StatusOK, gin.H{"response_content": fileID, "response_code": "000"})

}

//FileDownloadAPI : Download file from usager dossier
func FileDownloadAPI(c *gin.Context) {

	idExamen := c.Param("identifiant")
	fileID := c.Query("file_id")
	log.Println("File ID", fileID)
	examen, err := repository.FindUsagerExamenByIdentifiant(idExamen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response_content": "unkonwn-examen", "response_code": "100"})
		return
	}
	log.Println(examen)

	// Upload the file to gridfs
	file := utils.DownloadFile(fileID)
	im, str, error := image.Decode(bytes.NewReader(file))
	log.Println(str)
	log.Println(error)

	c.Header("Content-Disposition", "attachment; filename="+fileID+".jpeg")
	c.Header("Content-Type", "image/jpeg")
	c.Stream(func(w io.Writer) bool {
		jpeg.Encode(w, im, nil)
		return false
	})
	c.Writer.Write(file)

}
