package usager

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usager : regroupe l'ensemble des patients de la plateforme
type Usager struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Matricule             string             `bson:"matricule" unique:"true"`
	FirstName             string             `bson:"first_name"`
	LastName              string             `bson:"last_name"`
	DateOfBirth           time.Time          `bson:"date_of_birth"`
	PlaceOfBirth          string             `bson:"place_of_birth"`
	Address               string             `bson:"address"`
	PhoneNumber           string             `bson:"phone_number"`
	IdentityNumber        string             `bson:"identity_number"`
	TypeDocument          string             `bson:"type_document"`
	Sexe                  string             `bson:"sexe"`
	SituationMatrimoniale string             `bson:"situation_matrimoniale"`
	PersonneaPrevenir     PersonneaPrevenir  `bson:"personne_a_prevenir"`
	CreatedAt             time.Time          `bson:"created_at"`
}

// PersonneaPrevenir : Personne à prévenur en cas d'accident
type PersonneaPrevenir struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	FirstName          string             `bson:"first_name"`
	LastName           string             `bson:"last_name"`
	Address            string             `bson:"address"`
	PhoneNumber        string             `bson:"phone_number"`
	Sexe               string             `bson:"sexe"`
	RelationWithUsager string             `bson:"relation_with_usager"`
	CreatedAt          time.Time          `bson:"created_at"`
}
