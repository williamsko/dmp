package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddContenuHospitalisationToPatientRecord :  add consultation to dosser usager
func AddContenuHospitalisationToPatientRecord(patientRecord dossier.PatientRecord, hospitalisationPayload dossier.NewHostpitalisationPayloadValidator, agent entity.Agent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hospitalisationCollection := db.ConnectDb().Collection("hospitalisation")
	hospitalisation := &dossier.Hospitalisation{
		ID:                   primitive.NewObjectID(),
		Agent:                agent,
		PatientRecord:        patientRecord.ID,
		MotifHospitalisation: hospitalisationPayload.MotifHospitalisation,
	}
	_, err := hospitalisationCollection.InsertOne(ctx, hospitalisation)
	return err
}

// GetAllHospitalisationsByPatientRecord : Retreive all hospitalisations for usager
func GetAllHospitalisationsByPatientRecord(patientRecord *dossier.PatientRecord) ([]dossier.Hospitalisation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var hospitalisationUsager []dossier.Hospitalisation
	hospitalisationCollection := db.ConnectDb().Collection("consultation")
	cursor, err := hospitalisationCollection.Find(ctx, bson.M{"dossier": patientRecord.ID})
	if err = cursor.All(ctx, &hospitalisationUsager); err != nil {
		panic(err)
	}
	return hospitalisationUsager, err
}
