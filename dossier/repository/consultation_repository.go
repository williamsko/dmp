package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"go.mongodb.org/mongo-driver/bson"

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

// GetAllConsultationsByDossierUsager : Retreive all consultations for usager
func GetAllConsultationsByDossierUsager(dossierMedical *dossier.DossierMedical) ([]dossier.Consultation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var consultationsUsager []dossier.Consultation
	consultationCollection := db.ConnectDb().Collection("consultation")
	cursor, err := consultationCollection.Find(ctx, bson.M{"dossier": dossierMedical.ID})
	if err = cursor.All(ctx, &consultationsUsager); err != nil {
		panic(err)
	}
	return consultationsUsager, err
}
