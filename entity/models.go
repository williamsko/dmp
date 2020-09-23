package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	ID          primitive.ObjectID `json:"_id"`
	BrandNane   string             `json:"brand_name"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phone_number"`
	City        string             `json:"city"`
	CreatedAt   time.Time          `json:"created_at"`
}

// Agent : Doctor , nurse etc ...
type Agent struct {
	ID          primitive.ObjectID `json:"_id"`
	Entity      primitive.ObjectID `json:"entity"`
	Matricule   string             `json:"matricule"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phone_number"`
	City        string             `json:"city"`
	Profession  string             `json:"profession"`
	Specialite  string             `json:"specialite"`
	CreatedAt   time.Time          `json:"created_at"`
}

// AgentJob : Liste des types de documents
const (
	DOCTEUR    = "DOCTEUR"
	INFIRMIER  = "INFIRMIER"
	MATRONNE   = "MATRONNE"
	SECRETAIRE = "SECRETAIRE"
)
