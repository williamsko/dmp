package usager

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Usager : regroupe l'ensemble des usagers de la plateforme
type Usager struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Matricule      string             `bson:"matricule,omitempty"`
	FirstName      string             `bson:"first_name,omitempty"`
	LastName       string             `bson:"last_name,omitempty"`
	Address        string             `bson:"address,omitempty"`
	PhoneNumber    string             `bson:"phone_number,omitempty"`
	IdentityNumber string             `bson:"identity_number,omitempty"`
	TypeDocument   string             `bson:"typedocument,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}

// TypeDocument : Liste des types de documents
const (
	PASSEPORT = "PASSEPORT"
	CNI       = "CARTE_D_IDENTITE"
	CC        = "CARTE_CONSULAIRE"
	PC        = "PERMIS_DE_CONDUIRE"
)
