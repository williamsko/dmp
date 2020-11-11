package record

import (
	"bytes"
	"dmp/record"
	repository "dmp/record/repository"
	"dmp/entity"
	"dmp/patient"
	"dmp/utils"
	"image"
	jpeg "image/jpeg"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostExamenAPI : api to add new consultation to dossier
func PostExamenAPI(c *gin.Context) {
	var payload record.NewExamenValidator
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usagerMatricule := c.Param("matricule")
	patient, err := patient.FindPatientByMatricule(usagerMatricule)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-patient")
		return
	}
	agentMatricule, _ := c.Get("agent")
	agent, err := entity.FindAgentByMatricule(agentMatricule.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "unkonwn-agent")
		return
	}

	
	dossier, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "medical-record-does-not-exist")
		return
	}
	examen, err := repository.AddContenuExamenToPatientRecord(dossier, payload, agent)
	c.JSON(http.StatusOK, gin.H{"response_content": examen, "response_code": "000"})
}

//GetExamenAPI : api to get usager examen
func GetExamenAPI(c *gin.Context) {
	patient, err := patient.FindPatientByMatricule(c.Param("matricule"))
	patientRecord, err := repository.FindRecordByPatientD(patient.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "medical-record-does-not-exist")
		return
	}
	examensPatient, err := repository.GetAllExamensByPatientRecord(&patientRecord)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "dossier-creation-error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"response_content": examensPatient, "response_code": "000"})
}

//PatchExamenAPI : updateExamen with results
func PatchExamenAPI(c *gin.Context) {
	var payload record.UpdateExamenValidator
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
