package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"time"
)

// AddContenuHospitalisationUsagerToDossier :  add consultation to dosser usager
func AddContenuHospitalisationUsagerToDossier(dossierMedical dossier.DossierMedical, hospitalisationPayload dossier.NewHostpitalisationPayloadValidator, agent entity.Agent) (*dossier.Hospitalisation, error) {
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
	return hospitalisation, err

}
