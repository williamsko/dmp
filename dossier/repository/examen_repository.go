package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// AddContenuExamenUsagerToDossier :  add consultation to dosser usager
func AddContenuExamenUsagerToDossier(dossierMedical dossier.DossierMedical,
	examenPayload dossier.NewExamenValidator,
	agent entity.Agent) (*dossier.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	examenCollection := db.ConnectDb().Collection("examen")
	examen := &dossier.Examen{
		Agent:          agent.ID,
		Entity:         agent.Entity,
		DossierMedical: dossierMedical.ID,
		Type:           examenPayload.Type,
		Statut:         calculateExamenStatut(examenPayload.Content),
		Content:        examenPayload.Content,
	}
	_, err := examenCollection.InsertOne(ctx, examen)
	return examen, err
}

// GetAllExamensByDossierUsager : Retreive all examens for usager
func GetAllExamensByDossierUsager(dossierMedical *dossier.DossierMedical) ([]dossier.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var examensUsager []dossier.Examen
	examensCollection := db.ConnectDb().Collection("examen")
	cursor, err := examensCollection.Find(ctx, bson.M{"dossier": dossierMedical.ID})
	if err = cursor.All(ctx, &examensUsager); err != nil {
		panic(err)
	}
	return examensUsager, err
}

func calculateExamenStatut(content []dossier.ExamenContent) string {
	contentLen := len(content)
	if contentLen == 0 {
		return "EN_ATTENTE"
	}
	return "EFFECTUE"
}

// FindUsagerExamenByIdentifiant : Find usager examen for update
func FindUsagerExamenByIdentifiant(ID string) (dossier.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var examen dossier.Examen
	examenCollection := db.ConnectDb().Collection("examen")
	examenID, err := primitive.ObjectIDFromHex(ID)
	log.Print(examenID)
	log.Print(err)

	err = examenCollection.FindOne(ctx, bson.M{"_id": examenID}).Decode(&examen)
	return examen, err
}

// UpdateContenuExamen :  add consultation to dosser usager
func UpdateContenuExamen(examenPayload dossier.UpdateExamenValidator, examen dossier.Examen) (int64, error) {
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
func UpdateContenuExamenWithFile(examen dossier.Examen, fileID string) (int64, error) {
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

func createExamenContent(fileID string) dossier.ExamenContentFiles {
	examenContentFile := dossier.ExamenContentFiles{
		ID: fileID,
	}
	return examenContentFile
}
