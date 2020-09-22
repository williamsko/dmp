package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// AgentPayloadValidator : use to validate agent payload in all models
type AgentPayloadValidator struct {
	Matricule string `bson:"matricule,omitempty" binding:"required"`
}

// EntityPayloadValidator : use to validate entity payload in all models
type EntityPayloadValidator struct {
	Matricule string `bson:"matricule,omitempty" binding:"required"`
}

// Entity : Hospital for example
type Entity struct {
	ID          primitive.ObjectID `bson:"_id"`
	BrandNane   string             `bson:"brand_name"`
	Address     string             `bson:"address"`
	PhoneNumber string             `bson:"phone_number"`
	City        string             `bson:"city"`
	CreatedAt   time.Time          `bson:"created_at"`
}

// Agent : Doctor , nurse etc ...
type Agent struct {
	ID          primitive.ObjectID `bson:"_id"`
	Entity      primitive.ObjectID `bson:"entity"`
	Matricule   string             `bson:"matricule"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Address     string             `bson:"address"`
	PhoneNumber string             `bson:"phone_number"`
	City        string             `bson:"city"`
	Profession  string             `bson:"profession"`
	Specialite  string             `bson:"specialite"`
	CreatedAt   time.Time          `bson:"created_at"`
}

// AgentJob : Liste des types de documents
const (
	DOCTEUR    = "DOCTEUR"
	INFIRMIER  = "INFIRMIER"
	MATRONNE   = "MATRONNE"
	SECRETAIRE = "SECRETAIRE"
)
