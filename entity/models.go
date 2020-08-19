package usager

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Entity : Hospital for example
type Entity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BrandNane   string             `bson:"brand_name,omitempty"`
	Address     string             `bson:"address,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	City        string             `bson:"city,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// Agent : Doctor , nurse etc ...
type Agent struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Entity      primitive.ObjectID `bson:"entity,omitempty"`
	Matricule   string             `bson:"matricule,omitempty"`
	FirstName   string             `bson:"first_name,omitempty"`
	LastName    string             `bson:"last_name,omitempty"`
	Address     string             `bson:"address,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	City        string             `bson:"city,omitempty"`
	Profession  string             `bson:"profession,omitempty"`
	Specialite  string             `bson:"specialite,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// AgentJob : Liste des types de documents
const (
	DOCTEUR    = "DOCTEUR"
	INFIRMIER  = "INFIRMIER"
	MATRONNE   = "MATRONNE"
	SECRETAIRE = "SECRETAIRE"
)
