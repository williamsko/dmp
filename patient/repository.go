package patient

import (
	"context"
	"dmp/db"
	"dmp/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindPatientByPhoneNumber : Find usager by giving phone number
func FindPatientByPhoneNumber(phoneNumber string) (Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var patient Patient
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&patient)
	return patient, err
}

// FindPatientByMatricule : Find usager by giving matricule
func FindPatientByMatricule(matricule string) (Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var patient Patient
	usagerCollection := db.ConnectDb().Collection("usager")
	err := usagerCollection.FindOne(ctx, bson.M{"matricule": matricule}).Decode(&patient)
	return patient, err
}

// FindPatientByID : Find usager by giving matricule
func FindPatientByID(id primitive.ObjectID) (Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var patient Patient
	usagerCollection := db.ConnectDb().Collection("usager")

	err := usagerCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&patient)
	return patient, err
}

// CreateNewUsager : create a new usager
func CreateNewUsager(patient *Patient) (*Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	matricule := utils.GenerateRandomNumber()
	newPerson := createPersonneaPrevenir(patient.PersonneaPrevenir)
	newPatient := &Patient{
		ID:                    primitive.NewObjectID(),
		Matricule:             matricule,
		FirstName:             patient.FirstName,
		LastName:              patient.LastName,
		Address:               patient.Address,
		PhoneNumber:           patient.PhoneNumber,
		IdentityNumber:        patient.IdentityNumber,
		Sexe:                  patient.Sexe,
		TypeDocument:          patient.TypeDocument,
		SituationMatrimoniale: patient.SituationMatrimoniale,
		DateOfBirth:           patient.DateOfBirth,
		PersonneaPrevenir:     newPerson,
		CreatedAt:             time.Now(),
	}
	usagerCollection := db.ConnectDb().Collection("usager")
	_, err := usagerCollection.InsertOne(ctx, newPatient)

	return newPatient, err
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
func GetAllUsers() ([]Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var patients []Patient
	usagerCollection := db.ConnectDb().Collection("usager")
	cursor, err := usagerCollection.Find(ctx, bson.M{})
	if err = cursor.All(ctx, &patients); err != nil {
		panic(err)
	}
	return patients, err
}
