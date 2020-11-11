package record

import (
	"context"
	"dmp/db"
	"dmp/record"
	"dmp/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddContenuExamenToPatientRecord :  add consultation to dosser usager
func AddContenuExamenToPatientRecord(patientRecord record.PatientRecord,
	examenPayload record.NewExamenValidator,
	agent entity.Agent) (*record.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	examenCollection := db.ConnectDb().Collection("examen")
	examen := &record.Examen{
		ID:            primitive.NewObjectID(),
		Agent:         agent,
		Entity:        agent.Entity,
		PatientRecord: patientRecord.ID,
		Type:          examenPayload.Type,
		Statut:        calculateExamenStatut(examenPayload.Content),
		Content:       examenPayload.Content,
	}
	_, err := examenCollection.InsertOne(ctx, examen)
	return examen, err
}

// GetAllExamensByPatientRecord : Retreive all examens for usager
func GetAllExamensByPatientRecord(patientRecord *record.PatientRecord) ([]record.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var examensPatient []record.Examen
	examensCollection := db.ConnectDb().Collection("examen")
	cursor, err := examensCollection.Find(ctx, bson.M{"dossier": patientRecord.ID})
	if err = cursor.All(ctx, &examensPatient); err != nil {
		panic(err)
	}
	return examensPatient, err
}

func calculateExamenStatut(content []record.ExamenContent) string {
	contentLen := len(content)
	if contentLen == 0 {
		return "EN_ATTENTE"
	}
	return "EFFECTUE"
}

// FindUsagerExamenByIdentifiant : Find usager examen for update
func FindUsagerExamenByIdentifiant(ID string) (record.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var examen record.Examen
	examenCollection := db.ConnectDb().Collection("examen")
	examenID, err := primitive.ObjectIDFromHex(ID)
	log.Print(examenID)
	log.Print(err)

	err = examenCollection.FindOne(ctx, bson.M{"_id": examenID}).Decode(&examen)
	return examen, err
}

// UpdateContenuExamen :  add consultation to dosser usager
func UpdateContenuExamen(examenPayload record.UpdateExamenValidator, examen record.Examen) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	examenCollection := db.ConnectDb().Collection("examen")
	result, err := examenCollection.UpdateOne(
		ctx,
		bson.M{"_id": examen.ID},
		bson.D{
			{"$set", bson.D{
				{Key: "content", Value: examenPayload.Content},
				{Key: "statut", Value: calculateExamenStatut(examenPayload.Content)}},
			},
		},
	)
	return result.ModifiedCount, err
}

// UpdateContenuExamenWithFile :  add consultation to dosser usager
func UpdateContenuExamenWithFile(examen record.Examen, fileID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	examenCollection := db.ConnectDb().Collection("examen")
	examenContentFile := createExamenContent(fileID)
	examenContentFiles := append(examen.Files, examenContentFile)
	result, err := examenCollection.UpdateOne(
		ctx,
		bson.M{"_id": examen.ID},
		bson.D{
			{"$set", bson.D{{Key: "files", Value: examenContentFiles}}}})

	return result.ModifiedCount, err
}

func createExamenContent(fileID string) record.ExamenContentFiles {
	examenContentFile := record.ExamenContentFiles{
		ID: fileID,
	}
	return examenContentFile
}
