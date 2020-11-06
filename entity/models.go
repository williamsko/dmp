package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AgentPayloadValidator : use to validate agent payload in all models
type AgentPayloadValidator struct {
	Matricule string `json:"matricule,omitempty" binding:"required"`
}

// EntityPayloadValidator : use to validate entity payload in all models
type EntityPayloadValidator struct {
	Matricule string `json:"matricule,omitempty" binding:"required"`
}

// Entity : Hospital for example
type Entity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BrandNane   string             `bson:"brand_name"`
	Address     string             `bson:"address"`
	PhoneNumber string             `bson:"phone_number"`
	City        string             `bson:"city"`
	CreatedAt   time.Time          `bson:"created_at"`
}

// Agent : Doctor , nurse etc ...
type Agent struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Entity         primitive.ObjectID `bson:"entity" json:"entity"`
	Matricule      string             `bson:"matricule" json:"matricule"`
	Gender         string             `bson:"gender" json:"gender"`
	FirstName      string             `bson:"first_name" json:"first_name"`
	LastName       string             `bson:"last_name" json:"last_name"`
	Address        string             `bson:"address" json:"address"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number"`
	Email          string             `bson:"email" json:"email"`
	City           string             `bson:"city" json:"city"`
	Profession     string             `bson:"profession" json:"profession"`
	HashedPassword string             `bson:"hashed_password" json:"-"`
	DateOfBitrh    time.Time          `bson:"date_of_birth" json:"date_of_birth"`
	PlaceOfBirth   string             `bson:"place_of_birth" json:"place_of_birth"`
	CountryOfBirth string             `bson:"country_of_birth" json:"country_of_birth"`
	Nationality    string             `bson:"nationality" json:"nationality"`
	IDCardNumber   string             `bson:"id_card_number" json:"id_card_number"`
	MaritalStatus  string             `bson:"marital_status" json:"marital_status"`
	Specialite     string             `bson:"specialite" json:"specialite"`
	Service        string             `bson:"service" json:"service"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
}

// AgentJob : Liste des types de documents
const (
	DOCTEUR    = "DOCTEUR"
	INFIRMIER  = "INFIRMIER"
	MATRONNE   = "MATRONNE"
	SECRETAIRE = "SECRETAIRE"
)
