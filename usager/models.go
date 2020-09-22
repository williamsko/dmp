package usager

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Usager : regroupe l'ensemble des usagers de la plateforme
type Usager struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Matricule             string             `bson:"matricule" unique:"true"`
	FirstName             string             `bson:"first_name" binding:"required"`
	LastName              string             `bson:"last_name" binding:"required"`
	Address               string             `bson:"address" binding:"required"`
	PhoneNumber           string             `bson:"phone_number" binding:"required"`
	IdentityNumber        string             `bson:"identity_number" binding:"required"`
	TypeDocument          string             `bson:"type_document" binding:"required"`
	Sexe                  string             `bson:"sexe" binding:"required"`
	SituationMatrimoniale string             `bson:"situation_matrimoniale" binding:"required"`
	CreatedAt             time.Time          `bson:"created_at" bson:"created_at" time_format:"2006-01-02"`
}
