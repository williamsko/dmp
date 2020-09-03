package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"go.mongodb.org/mongo-driver/bson"

	"time"
)

// AddContenuExamenUsagerToDossier :  add consultation to dosser usager
func AddContenuExamenUsagerToDossier(dossierMedical dossier.DossierMedical, examenPayload dossier.NewExamenValidator, agent entity.Agent) (*dossier.Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	examenCollection := db.ConnectDb().Collection("examen")

	examen := &dossier.Examen{
		Agent:          agent.ID,
		Entity:         agent.Entity,
		DossierMedical: dossierMedical.ID,
		Type:           examenPayload.Type,
		Content:        examenPayload.Content,
	}
	_, err := examenCollection.InsertOne(ctx, examen)
	return examen, err

}

// GetAllExamensByDossierUsager : Retreive all examens for usager
func GetAllExamensByDossierUsager(dossierMedical *dossier.DossierMedical) ([]dossier.Consultation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var examensUsager []dossier.Consultation
	examensCollection := db.ConnectDb().Collection("examen")
	cursor, err := examensCollection.Find(ctx, bson.M{"dossier": dossierMedical.ID})
	if err = cursor.All(ctx, &examensUsager); err != nil {
		panic(err)
	}
	return examensUsager, err
}
