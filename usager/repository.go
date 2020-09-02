package usager

import (
	"context"
	"dmp/db"
	"dmp/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// FindUsagerByPhoneNumber : Find usager
func FindUsagerByPhoneNumber(phoneNumber string) (Usager, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"phonenumber": phoneNumber}).Decode(&usager)
	return usager, err
}

// FindUsagerByMatricule : Find usager
func FindUsagerByMatricule(matricule string) (Usager, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"matricule": matricule}).Decode(&usager)
	return usager, err
}

// CreateNewUsager : create a new usager
func CreateNewUsager(usager *NewUsagerPayloadValidator) (*Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	matricule := utils.RandomObjectMatricule(10)
	newUsager := &Usager{
		Matricule:             matricule,
		FirstName:             usager.FirstName,
		LastName:              usager.LastName,
		Address:               usager.Address,
		PhoneNumber:           usager.PhoneNumber,
		IdentityNumber:        usager.IdentityNumber,
		Sexe:                  usager.Sexe,
		SituationMatrimoniale: usager.SituationMatrimoniale,
	}
	fmt.Println(matricule)
	usagerCollection := db.ConnectDb().Collection("usager")
	_, err := usagerCollection.InsertOne(ctx, newUsager)
	return newUsager, err
}
