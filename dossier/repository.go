package dossier

import (
	"context"
	"dmp/db"
	"dmp/entity"
	"dmp/usager"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"dmp/utils"
	"time"
)

// FindDossierByUsagerID : Find usager dossier
func FindDossierByUsagerID(_id primitive.ObjectID) (Dossier, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dossier := Dossier{}

	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"usager": _id}).Decode(&dossier)
	return dossier, err
}

// CreateEmptyDossier : create a new usager
func CreateEmptyDossier(usager usager.Usager, agent entity.Agent) (*Dossier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dossierCollection := db.ConnectDb().Collection("dossier")
	numberDossier := utils.RandomObjectMatricule(10)

	dossier := &Dossier{
		Usager: usager.ID,
		Agent:  agent.ID,
		Entity: agent.Entity,
		Number: numberDossier,
	}
	_, err := dossierCollection.InsertOne(ctx, dossier)
	return dossier, err
}

// AddContenuAntecedentUsagerToDossier :  add antecedent to dosser usager
func AddContenuAntecedentUsagerToDossier(dossier Dossier, antecedentPayload NewAntecedentPayloadValidator, agent entity.Agent) (*Antecedent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	antecedentCollection := db.ConnectDb().Collection("antecedent")

	antecedent := &Antecedent{
		Agent:                 agent.ID,
		Entity:                agent.Entity,
		Dossier:               dossier.ID,
		AntecedentMedical:     antecedentPayload.AntecedentMedical,
		AntecedentChirurgical: antecedentPayload.AntecedentChirurgical,
		AntecedentFamilial:    antecedentPayload.AntecedentFamilial,
		ModeDeVie:             antecedentPayload.ModeDeVie,
	}
	_, err := antecedentCollection.InsertOne(ctx, antecedent)
	fmt.Println(err)
	return antecedent, err

}

// AddContenuConsultationUsagerToDossier :  add consultation to dosser usager
func AddContenuConsultationUsagerToDossier(dossier Dossier, consultationPayload NewConsultationPayloadValidator, agent entity.Agent) (*Consultation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	consultationCollection := db.ConnectDb().Collection("consultation")

	consultation := &Consultation{
		Agent:             agent.ID,
		Entity:            agent.Entity,
		Dossier:           dossier.ID,
		HistoireMaladie:   consultationPayload.HistoireMaladie,
		MotifConsultation: consultationPayload.MotifConsultation,
		Commentaire:       consultationPayload.Commentaire,
	}
	_, err := consultationCollection.InsertOne(ctx, consultation)
	return consultation, err

}

// AddContenuHospitalisationUsagerToDossier :  add consultation to dosser usager
func AddContenuHospitalisationUsagerToDossier(dossier Dossier, hospitalisationPayload NewHostpitalisationPayloadValidator, agent entity.Agent) (*Hospitalisation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hospitalisationCollection := db.ConnectDb().Collection("hospitalisation")

	hospitalisation := &Hospitalisation{
		Agent:                agent.ID,
		Entity:               agent.Entity,
		Dossier:              dossier.ID,
		MotifHospitalisation: hospitalisationPayload.MotifHospitalisation,
		Commentaire:          hospitalisationPayload.Commentaire,
	}
	_, err := hospitalisationCollection.InsertOne(ctx, hospitalisation)
	return hospitalisation, err

}

// AddContenuExamenUsagerToDossier :  add consultation to dosser usager
func AddContenuExamenUsagerToDossier(dossier Dossier, examenPayload NewExamenValidator, agent entity.Agent) (*Examen, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	examenCollection := db.ConnectDb().Collection("examen")

	examen := &Examen{
		Agent:   agent.ID,
		Entity:  agent.Entity,
		Dossier: dossier.ID,
		Type:    examenPayload.Type,
		Content: examenPayload.Content,
	}
	_, err := examenCollection.InsertOne(ctx, examen)
	return examen, err

}
