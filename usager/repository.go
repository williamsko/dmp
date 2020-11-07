package usager

import (
	"context"
	"dmp/db"
	"dmp/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindUsagerByPhoneNumber : Find usager by giving phone number
func FindUsagerByPhoneNumber(phoneNumber string) (Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&usager)
	return usager, err
}

// FindUsagerByMatricule : Find usager by giving matricule
func FindUsagerByMatricule(matricule string) (Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"matricule": matricule}).Decode(&usager)
	return usager, err
}

// FindUsagerByID : Find usager by giving matricule
func FindUsagerByID(id primitive.ObjectID) (Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var usager Usager
	usagerCollection := db.ConnectDb().Collection("usager")

	err := usagerCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&usager)
	return usager, err
}

// CreateNewUsager : create a new usager
func CreateNewUsager(usager *Usager) (*Usager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	matricule := utils.GenerateRandomNumber()
	newPerson := createPersonneaPrevenir(usager.PersonneaPrevenir)
	newUsager := &Usager{
		ID:                    primitive.NewObjectID(),
		Matricule:             matricule,
		FirstName:             usager.FirstName,
		LastName:              usager.LastName,
		Address:               usager.Address,
		PhoneNumber:           usager.PhoneNumber,
		IdentityNumber:        usager.IdentityNumber,
		Sexe:                  usager.Sexe,
		TypeDocument:          usager.TypeDocument,
		SituationMatrimoniale: usager.SituationMatrimoniale,
		DateOfBirth:           usager.DateOfBirth,
		PersonneaPrevenir:     newPerson,
		CreatedAt:             time.Now(),
	}
	usagerCollection := db.ConnectDb().Collection("usager")
	_, err := usagerCollection.InsertOne(ctx, newUsager)

	return newUsager, err
}

// createPersonneaPrevenir : create personne a prevenir
func createPersonneaPrevenir(personneaPrevenir PersonneaPrevenir) PersonneaPrevenir {
	newPerson := PersonneaPrevenir{
		ID:                 primitive.NewObjectID(),
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
