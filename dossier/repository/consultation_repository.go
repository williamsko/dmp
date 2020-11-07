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
func AddContenuConsultationUsagerToDossier(patientRecord dossier.PatientRecord,
	consultationPayload dossier.NewConsultationPayloadValidator, agent entity.Agent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	consultationCollection := db.ConnectDb().Collection("consultation")
	consultation := &dossier.Consultation{
		Agent:             agent.ID,
		Entity:            agent.Entity.ID,
		PatientRecord:    patientRecord.ID,
		HistoireMaladie:   consultationPayload.HistoireMaladie,
		MotifConsultation: consultationPayload.MotifConsultation,
		CreatedAt:         time.Now(),
	}
	_, err := consultationCollection.InsertOne(ctx, consultation)
	return err

}

// GetAllConsultationsByDossierUsager : Retreive all consultations for usager
func GetAllConsultationsByDossierUsager(patientRecord *dossier.PatientRecord) ([]dossier.Consultation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var consultationsUsager []dossier.Consultation
	consultationCollection := db.ConnectDb().Collection("consultation")
	cursor, err := consultationCollection.Find(ctx, bson.M{"dossier": patientRecord.ID})
	if err = cursor.All(ctx, &consultationsUsager); err != nil {
		panic(err)
	}
	return consultationsUsager, err
}
