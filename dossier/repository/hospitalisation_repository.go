package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddContenuHospitalisationUsagerToDossier :  add consultation to dosser usager
func AddContenuHospitalisationUsagerToDossier(dossierMedical dossier.DossierMedical, hospitalisationPayload dossier.NewHostpitalisationPayloadValidator, agent entity.Agent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hospitalisationCollection := db.ConnectDb().Collection("hospitalisation")
	hospitalisation := &dossier.Hospitalisation{
		Agent:                agent.ID,
		Entity:               agent.Entity,
		DossierMedical:       dossierMedical.ID,
		MotifHospitalisation: hospitalisationPayload.MotifHospitalisation,
		Commentaire:          hospitalisationPayload.Commentaire,
	}
	_, err := hospitalisationCollection.InsertOne(ctx, hospitalisation)
	return err
}

// GetAllHospitalisationsByDossierUsager : Retreive all hospitalisations for usager
func GetAllHospitalisationsByDossierUsager(dossierMedical *dossier.DossierMedical) ([]dossier.Hospitalisation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var hospitalisationUsager []dossier.Hospitalisation
	hospitalisationCollection := db.ConnectDb().Collection("consultation")
	cursor, err := hospitalisationCollection.Find(ctx, bson.M{"dossier": dossierMedical.ID})
	if err = cursor.All(ctx, &hospitalisationUsager); err != nil {
		panic(err)
	}
	return hospitalisationUsager, err
}
