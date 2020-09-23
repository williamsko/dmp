package usager

import (
	"context"
	"dmp/db"
	"dmp/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUsagerByPhoneNumber : Find usager
func FindUsagerByPhoneNumber(phoneNumber string) (Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&usager)
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
	matricule := utils.GenerateRandomNumber()
	newPerson := createPersonneaPrevenir(usager.PersonneaPrevenirValidator)
	newUsager := &Usager{
		Matricule:             matricule,
		FirstName:             usager.FirstName,
		LastName:              usager.LastName,
		Address:               usager.Address,
		PhoneNumber:           usager.PhoneNumber,
		IdentityNumber:        usager.IdentityNumber,
		Sexe:                  usager.Sexe,
		TypeDocument:          usager.TypeDocument,
		SituationMatrimoniale: usager.SituationMatrimoniale,
		PersonneaPrevenir:     newPerson,
		CreatedAt:             time.Now(),
	}
	usagerCollection := db.ConnectDb().Collection("usager")
	_, err := usagerCollection.InsertOne(ctx, newUsager)

	return newUsager, err
}

// createPersonneaPrevenir : create personne a prevenir
func createPersonneaPrevenir(personneaPrevenir PersonneaPrevenirValidator) PersonneaPrevenir {
	newPerson := PersonneaPrevenir{
		FirstName:          personneaPrevenir.FirstName,
		LastName:           personneaPrevenir.LastName,
		Address:            personneaPrevenir.Address,
		PhoneNumber:        personneaPrevenir.PhoneNumber,
		Sexe:               personneaPrevenir.Sexe,
		RelationWithUsager: personneaPrevenir.RelationWithUsager,
		CreatedAt:          time.Now(),
	}
	return newPerson
}

// GetAllUsers : Retreive all usager
func GetAllUsers() ([]Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usagers []Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	cursor, err := usagerCollection.Find(ctx, bson.M{})
	if err = cursor.All(ctx, &usagers); err != nil {
		panic(err)
	}
	return usagers, err
}
