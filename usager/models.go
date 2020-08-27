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
	LastName       string             `json:"last_name,omitempty" binding:"required"`
	Address        string             `json:"address,omitempty" binding:"required"`
	PhoneNumber    string             `json:"phone_number,omitempty" binding:"required"`
	IdentityNumber string             `json:"identity_number,omitempty" binding:"required"`
	TypeDocument   string             `json:"type_document,omitempty" binding:"required"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at" time_format:"2006-01-02"`
}

// TypeDocument : Liste des types de documents
const (
	PASSEPORT = "PASSEPORT"
	CNI       = "CARTE_D_IDENTITE"
	CC        = "CARTE_CONSULAIRE"
	PC        = "PERMIS_DE_CONDUIRE"
)
