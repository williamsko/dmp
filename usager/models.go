package usager

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Usager : regroupe l'ensemble des usagers de la plateforme
type Usager struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Matricule      string             `json:"matricule,omitempty" unique:"true"`
	FirstName      string             `json:"first_name,omitempty" binding:"required"`
	LastName       string             `json:"last_name,omitempty"`
	Address        string             `json:"address,omitempty"`
	PhoneNumber    string             `json:"phone_number,omitempty"`
	IdentityNumber string             `json:"identity_number,omitempty"`
	TypeDocument   string             `json:"typedocument,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at" time_format:"2006-01-02"`
}

// TypeDocument : Liste des types de documents
const (
	PASSEPORT = "PASSEPORT"
	CNI       = "CARTE_D_IDENTITE"
	CC        = "CARTE_CONSULAIRE"
	PC        = "PERMIS_DE_CONDUIRE"
)
