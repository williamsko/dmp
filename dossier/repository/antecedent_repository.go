package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// AddContenuAntecedentUsagerToDossier :  add antecedent to dosser usager
func AddContenuAntecedentUsagerToDossier(dossierMedical dossier.DossierMedical, antecedentPayload dossier.NewAntecedentPayloadValidator, agent entity.Agent) (*dossier.Antecedent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	antecedentCollection := db.ConnectDb().Collection("antecedent")

	antecedent := &dossier.Antecedent{
		Agent:                 agent.ID,
		Entity:                agent.Entity,
		DossierMedical:        dossierMedical.ID,
		AntecedentMedical:     antecedentPayload.AntecedentMedical,
		AntecedentChirurgical: antecedentPayload.AntecedentChirurgical,
		AntecedentFamilial:    antecedentPayload.AntecedentFamilial,
		ModeDeVie:             antecedentPayload.ModeDeVie,
	}
	_, err := antecedentCollection.InsertOne(ctx, antecedent)
	fmt.Println(err)
	return antecedent, err

}

// GetAllAntecedentByDossierUsager : Retreive all antecedents for usager
func GetAllAntecedentByDossierUsager(dossierMedical *dossier.DossierMedical) ([]dossier.Antecedent, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var antedecentsUsager []dossier.Antecedent
	antedecentCollection := db.ConnectDb().Collection("antecedent")
	cursor, err := antedecentCollection.Find(ctx, bson.M{"dossier": dossierMedical.ID})
	if err = cursor.All(ctx, &antedecentsUsager); err != nil {
		panic(err)
	}
	return antedecentsUsager, err
}
