package dossier

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DossierMedical : DMP  for users
type DossierMedical struct {
	ID         primitive.ObjectID `json:"_id"`
	Usager     primitive.ObjectID `json:"usager"`
	Number     string             `json:"matricule"`
	CreatedAt  time.Time          `json:"created_at"`
	LastAccess time.Time          `json:"last_access"`
	Agent      primitive.ObjectID `json:"agent"`  //Agent who created the DMP
	Entity     primitive.ObjectID `json:"entity"` //In which entity the DMP is created
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
	CreatedAt      time.Time            `bson:"created_at"`
}

// Antecedent : Antecedent
type Antecedent struct {
	ID                    primitive.ObjectID `bson:"_id"`
	DossierMedical        primitive.ObjectID `bson:"dossier"`
	Agent                 primitive.ObjectID `bson:"agent"`
	Entity                primitive.ObjectID `bson:"entity"`
	AntecedentMedical     string             `bson:"antecedent_medical"`
	AntecedentChirurgical string             `bson:"antecedent_chirurgical"`
	AntecedentFamilial    string             `bson:"antecedent_familial"`
	ModeDeVie             string             `bson:"mode_de_vie"`
	CreatedAt             time.Time          `bson:"created_at"`
}

// Consultation : all consultations model
type Consultation struct {
	ID                primitive.ObjectID `bson:"_id"`
	DossierMedical    primitive.ObjectID `bson:"dossier"`
	Agent             primitive.ObjectID `bson:"agent"`
	Entity            primitive.ObjectID `bson:"entity"`
	MotifConsultation string             `bson:"motif_consultation"`
	HistoireMaladie   string             `bson:"histoire_maladie"`
	CreatedAt         time.Time          `bson:"created_at"`
}

// Hospitalisation : all Hospitalisation model
type Hospitalisation struct {
	ID                   primitive.ObjectID `bson:"_id"`
	DossierMedical       primitive.ObjectID `bson:"dossier"`
	Agent                primitive.ObjectID `bson:"agent"`
	Entity               primitive.ObjectID `bson:"entity"`
	MotifHospitalisation string             `bson:"motif_hospitalisation"`
	Commentaire          string             `bson:"commentaire"`
	CreatedAt            time.Time          `bson:"created_at"`
}

// ExamenContent : Examen content
type ExamenContent struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Value     string             `bson:"value"`
	CreatedAt time.Time          `bson:"created_at"`
}

// ExamenContentFiles : Examen content files
type ExamenContentFiles struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
}
