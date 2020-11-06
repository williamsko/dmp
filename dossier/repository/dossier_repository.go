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

// FindDossierByUsagerID : Find usager dossier
func FindDossierByUsagerID(_id primitive.ObjectID) (dossier.DossierMedical, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dossierMedical := dossier.DossierMedical{}
	dossierCollection := db.ConnectDb().Collection("dossier")
	err := dossierCollection.FindOne(ctx, bson.M{"usager": _id}).Decode(&dossierMedical)
	return dossierMedical, err
}

// CreateEmptyDossier : create a new usager
func CreateEmptyDossier(usager usager.Usager, medecinTraitant entity.Agent, agent entity.Agent) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dossierCollection := db.ConnectDb().Collection("dossier")
	numberDossier := utils.GenerateRandomNumber()
	dossierMedical := &dossier.DossierMedical{
		Usager:          usager.ID,
		MedecinTraitant: medecinTraitant.ID,
		Entity:          agent.Entity.ID,
		Number:          numberDossier,
		CreatedAt:       time.Now(),
	}
	_, err := dossierCollection.InsertOne(ctx, dossierMedical)
	return numberDossier, err
}
