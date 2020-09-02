package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"time"
)

// AddContenuConsultationUsagerToDossier :  add consultation to dosser usager
func AddContenuConsultationUsagerToDossier(dossierMedical dossier.DossierMedical, consultationPayload dossier.NewConsultationPayloadValidator, agent entity.Agent) (*dossier.Consultation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	consultationCollection := db.ConnectDb().Collection("consultation")

	consultation := &dossier.Consultation{
		Agent:             agent.ID,
		Entity:            agent.Entity,
		DossierMedical:    dossierMedical.ID,
		HistoireMaladie:   consultationPayload.HistoireMaladie,
		MotifConsultation: consultationPayload.MotifConsultation,
		Commentaire:       consultationPayload.Commentaire,
	}
	_, err := consultationCollection.InsertOne(ctx, consultation)
	return consultation, err

}
