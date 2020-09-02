package dossier

import (
	"context"
	"dmp/db"
	"dmp/dossier"
	"dmp/entity"
	"fmt"
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
