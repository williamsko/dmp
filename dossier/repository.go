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
func AddContenuAntecedentUsagerToDossier(dossier Dossier, antecedentPayload NewAntecedentPayload, agent entity.Agent) (*Antecedent, error) {
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
