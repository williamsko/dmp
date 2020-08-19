package usager

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Dossier : DMP  for users
type Dossier struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Number    string             `bson:"matricule,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Agent     primitive.ObjectID `bson:"agent,omitempty"`  //Agent who created the DMP
	Entity    primitive.ObjectID `bson:"entity,omitempty"` //In which entity the DMP is created
}

// ContenuDossier : DMP Content
type ContenuDossier struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Dossier     primitive.ObjectID `bson:"dossier,omitempty"`
	Agent       primitive.ObjectID `bson:"agent,omitempty"`
	Entity      primitive.ObjectID `bson:"entity,omitempty"`
	Content     string             `bson:"content,omitempty"`
	ContentType string             `bson:"content_type,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// TypeDocument : Liste des types de documents
const (
	ExamenGeneral              = "Examen Général"
	ExamenCardioVasculaire     = "Examen Cardio vasculaire"
	ExamenAppareilRespiratoire = "Examen de l'appareil respiratoire"
)
