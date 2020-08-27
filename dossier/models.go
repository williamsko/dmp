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
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Dossier        primitive.ObjectID `bson:"dossier,omitempty"`
	Agent          primitive.ObjectID `bson:"agent,omitempty"`
	Entity         primitive.ObjectID `bson:"entity,omitempty"`
	TypeAntecedent TypeAntecedent     `bson:"type_antecedent,omitempty"`
	Detail         string             `bson:"content_type,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}

// TypeExamen : Liste examens médicaux
type TypeExamen struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Code      string             `bson:"content_type,omitempty"`
	Detail    string             `bson:"content_type,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

//Type Antecedent
type TypeAntecedent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Tpy       string             `bson:"type,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

const (
	AntecedentMedical     = "Antecedent médical"
	AntecedentFamilial    = "Antecedent familial"
	AntecedentChirurgical = "Antecedent chirurgical"
	ModeDeVie             = "Mod de vie"
)

//DossierContentType
const (
	EXAMEN       = "EXAMEN"
	RESULTAT     = "RESULTAT"
	CONSULTATION = "CONSULTATION"
)
