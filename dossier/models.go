package dossier

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// DossierMedical : DMP  for users
type DossierMedical struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Usager     primitive.ObjectID `bson:"usager,omitempty"`
	Number     string             `bson:"matricule,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	LastAccess time.Time          `json:"last_access" bson:"last_access"`
	Agent      primitive.ObjectID `bson:"agent,omitempty"`  //Agent who created the DMP
	Entity     primitive.ObjectID `bson:"entity,omitempty"` //In which entity the DMP is created
}

// Examen : DMP Content
type Examen struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	DossierMedical primitive.ObjectID   `bson:"dossier,omitempty"`
	Agent          primitive.ObjectID   `bson:"agent,omitempty"`
	Entity         primitive.ObjectID   `bson:"entity,omitempty"`
	Content        []ExamenContent      `bson:"content,omitempty"`
	Files          []ExamenContentFiles `bson:"files,omitempty"`
	Statut         string               `bson:"statut,omitempty"`
	Type           string               `bson:"examen_type,omitempty"`
	CreatedAt      time.Time            `json:"created_at" bson:"created_at"`
}

// Antecedent : Antecedent
type Antecedent struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	DossierMedical        primitive.ObjectID `bson:"dossier,omitempty"`
	Agent                 primitive.ObjectID `bson:"agent,omitempty"`
	Entity                primitive.ObjectID `bson:"entity,omitempty"`
	AntecedentMedical     string             `bson:"antecedent_medical,omitempty"`
	AntecedentChirurgical string             `bson:"antecedent_chirurgical,omitempty"`
	AntecedentFamilial    string             `bson:"antecedent_familial,omitempty"`
	ModeDeVie             string             `bson:"mode_de_vie,omitempty"`
	CreatedAt             time.Time          `json:"created_at" bson:"created_at"`
}

// Consultation : all consultations model
type Consultation struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	DossierMedical    primitive.ObjectID `bson:"dossier,omitempty"`
	Agent             primitive.ObjectID `bson:"agent,omitempty"`
	Entity            primitive.ObjectID `bson:"entity,omitempty"`
	MotifConsultation string             `bson:"motif_consultation,omitempty"`
	HistoireMaladie   string             `bson:"histoire_maladie,omitempty"`
	Commentaire       string             `bson:"commentaire,omitempty"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
}

// Hospitalisation : all Hospitalisation model
type Hospitalisation struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	DossierMedical       primitive.ObjectID `bson:"dossier,omitempty"`
	Agent                primitive.ObjectID `bson:"agent,omitempty"`
	Entity               primitive.ObjectID `bson:"entity,omitempty"`
	MotifHospitalisation string             `bson:"motif_hospitalisation,omitempty"`
	Commentaire          string             `bson:"commentaire,omitempty"`
	CreatedAt            time.Time          `json:"created_at" bson:"created_at"`
}

// ExamenContent : Examen content
type ExamenContent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Value     string             `bson:"value,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// ExamenContentFiles : Examen content files
type ExamenContentFiles struct {
	ID        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
