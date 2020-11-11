package record

import (
	"context"
	"dmp/db"
	"dmp/record"

	"dmp/entity"
	"dmp/patient"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"dmp/utils"
	"time"
)

// FindRecordByPatientD : Find usager dossier
func FindRecordByPatientD(_id primitive.ObjectID) (record.PatientRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	patientRecord := record.PatientRecord{}
	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"usager": _id}).Decode(&patientRecord)
	return patientRecord, err
}

// CreateEmptyPatientRecord : create a new usager
func CreateEmptyPatientRecord(patient patient.Patient, medecinTraitant entity.Agent, agent entity.Agent) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dossierCollection := db.ConnectDb().Collection("dossier")
	numberDossier := utils.GenerateRandomNumber()
	patientRecord := &record.PatientRecord{
		Usager:          patient.ID,
		MedecinTraitant: medecinTraitant.ID,
		Entity:          agent.Entity.ID,
		Number:          numberDossier,
		CreatedAt:       time.Now(),
	}
	_, err := dossierCollection.InsertOne(ctx, patientRecord)
	return numberDossier, err
}

// FindPatientRecordByNumber : Find usager dossier
func FindPatientRecordByNumber(number string) (record.PatientRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	patientRecord := record.PatientRecord{}
	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"patient_medical_record_number": number}).Decode(&patientRecord)
	return patientRecord, err
}
