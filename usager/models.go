package usager

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usager : regroupe l'ensemble des patients de la plateforme
type Usager struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Matricule             string             `bson:"matricule" unique:"true" json:"matricule"`
	FirstName             string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName              string             `bson:"last_name" json:"last_name" binding:"required"`
	DateOfBirth           string             `bson:"date_of_birth" json:"date_of_birth" binding:"required"`
	PlaceOfBirth          string             `bson:"place_of_birth" json:"place_of_birth" binding:"required"`
	Address               string             `bson:"address" json:"address" binding:"required"`
	PhoneNumber           string             `bson:"phone_number" json:"phone_number" binding:"required"`
	IdentityNumber        string             `bson:"identity_number" json:"identity_number" binding:"required"`
	TypeDocument          string             `bson:"type_document" json:"type_document" binding:"required"`
	Sexe                  string             `bson:"sexe" json:"sexe" binding:"required"`
	SituationMatrimoniale string             `bson:"situation_matrimoniale" json:"situation_matrimoniale" binding:"required"`
	PersonneaPrevenir     PersonneaPrevenir  `bson:"personne_a_prevenir" json:"personne_a_prevenir" binding:"required"`
	CreatedAt             time.Time          `bson:"created_at" json:"created_at"`
}

// PersonneaPrevenir : Personne à prévenur en cas d'accident
type PersonneaPrevenir struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName          string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName           string             `bson:"last_name" json:"last_name" binding:"required"`
	Address            string             `bson:"address" json:"address" binding:"required"`
	PhoneNumber        string             `bson:"phone_number" json:"phone_number" binding:"required"`
	Sexe               string             `bson:"sexe" json:"sexe" binding:"required"`
	RelationWithUsager string             `bson:"relation_with_usager" json:"relation_with_usager" binding:"required"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at"`
}
