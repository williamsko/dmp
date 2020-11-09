package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddAntecedentToPatientRecord :  add antecedent to dosser usager
func AddAntecedentToPatientRecord(patientRecord dossier.PatientRecord, antecedentPayload dossier.NewAntecedentPayloadValidator, agent entity.Agent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	antecedentCollection := db.ConnectDb().Collection("antecedent")
	antecedent := &dossier.Antecedent{
		Agent:                 agent,
		Entity:                agent.Entity,
		PatientRecord:         patientRecord.ID,
		AntecedentMedical:     antecedentPayload.AntecedentMedical,
		AntecedentChirurgical: antecedentPayload.AntecedentChirurgical,
		AntecedentFamilial:    antecedentPayload.AntecedentFamilial,
		ModeDeVie:             antecedentPayload.ModeDeVie,
		CreatedAt:             time.Now(),
	}
	_, err := antecedentCollection.InsertOne(ctx, antecedent)
	return err
}

// GetAllAntecedentByPatientRecord : Retreive all antecedents for usager
func GetAllAntecedentByPatientRecord(patientRecord *dossier.PatientRecord) ([]dossier.Antecedent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var antedecentsUsager []dossier.Antecedent
	antedecentCollection := db.ConnectDb().Collection("antecedent")
	cursor, err := antedecentCollection.Find(ctx, bson.M{"dossier": patientRecord.ID})
	if err = cursor.All(ctx, &antedecentsUsager); err != nil {
		panic(err)
	}
	return antedecentsUsager, err
}
