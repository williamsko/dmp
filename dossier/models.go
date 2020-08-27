package dossier

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Dossier : DMP  for users
type Dossier struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Usager     primitive.ObjectID `bson:"usager,omitempty"`
	Number     string             `bson:"matricule,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	LastAccess time.Time          `json:"last_access" bson:"last_access"`
	Agent      primitive.ObjectID `bson:"agent,omitempty"`  //Agent who created the DMP
	Entity     primitive.ObjectID `bson:"entity,omitempty"` //In which entity the DMP is created
}

// ContenuDossier : DMP Content
type ContenuDossier struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Dossier     Dossier            `bson:"dossier,omitempty"`
	Agent       primitive.ObjectID `bson:"agent,omitempty"`
	Entity      primitive.ObjectID `bson:"entity,omitempty"`
	Content     string             `bson:"content,omitempty"`
	ContentType string             `bson:"content_type,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// Antecedent : Antecedent
type Antecedent struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Dossier               primitive.ObjectID `bson:"dossier,omitempty"`
	Agent                 primitive.ObjectID `bson:"agent,omitempty"`
	Entity                primitive.ObjectID `bson:"entity,omitempty"`
	AntecedentMedical     string             `bson:"antecedent_medical,omitempty"`
	AntecedentChirurgical string             `bson:"antecedent_chirurgical,omitempty"`
	AntecedentFamilial    string             `bson:"antecedent_familial,omitempty"`
	ModeDeVie             string             `bson:"mode_de_vie,omitempty"`
	CreatedAt             time.Time          `json:"created_at" bson:"created_at"`
}

//DossierContentType
const (
	EXAMEN       = "EXAMEN"
	RESULTAT     = "RESULTAT"
	CONSULTATION = "CONSULTATION"
)
