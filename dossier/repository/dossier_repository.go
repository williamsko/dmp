package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"

	"dmp/entity"
	"dmp/usager"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"dmp/utils"
	"time"
)

// FindPatientRecordByUsagerID : Find usager dossier
func FindPatientRecordByUsagerID(_id primitive.ObjectID) (dossier.PatientRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	patientRecord := dossier.PatientRecord{}
	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"usager": _id}).Decode(&patientRecord)
	return patientRecord, err
}

// CreateEmptyPatientRecord : create a new usager
func CreateEmptyPatientRecord(usager usager.Usager, medecinTraitant entity.Agent, agent entity.Agent) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dossierCollection := db.ConnectDb().Collection("dossier")
	numberDossier := utils.GenerateRandomNumber()
	patientRecord := &dossier.PatientRecord{
		Usager:          usager.ID,
		MedecinTraitant: medecinTraitant.ID,
		Entity:          agent.Entity.ID,
		Number:          numberDossier,
		CreatedAt:       time.Now(),
	}
	_, err := dossierCollection.InsertOne(ctx, patientRecord)
	return numberDossier, err
}

// FindPatientRecordByNumber : Find usager dossier
func FindPatientRecordByNumber(number string) (dossier.PatientRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	patientRecord := dossier.PatientRecord{}
	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"patient_medical_record_number": number}).Decode(&patientRecord)
	return patientRecord, err
}
