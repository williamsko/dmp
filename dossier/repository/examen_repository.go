package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
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
