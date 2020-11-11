package record

import (
	"context"
	"dmp/db"
	"dmp/entity"
	"dmp/record"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddAntecedentToPatientRecord :  add antecedent to dosser usager
func AddAntecedentToPatientRecord(patientRecord record.PatientRecord, antecedentPayload record.NewAntecedentPayloadValidator, agent entity.Agent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	antecedentCollection := db.ConnectDb().Collection("antecedent")
	antecedent := &record.Antecedent{
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

// GetAntecedentsByPatientRecord : Retreive all antecedents for usager
func GetAntecedentsByPatientRecord(patientRecord *record.PatientRecord) ([]record.Antecedent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var antedecentsPatient []record.Antecedent
	antedecentCollection := db.ConnectDb().Collection("antecedent")
	cursor, err := antedecentCollection.Find(ctx, bson.M{"dossier": patientRecord.ID})
	if err = cursor.All(ctx, &antedecentsPatient); err != nil {
		panic(err)
	}
	return antedecentsPatient, err
}
