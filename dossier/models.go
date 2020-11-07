package dossier

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PatientRecord : DMP  for users
type PatientRecord struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Usager          primitive.ObjectID `bson:"usager" json:"usager"`
	Number          string             `bson:"patient_medical_record_number" json:"patient_medical_record_number"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	LastAccess      time.Time          `bson:"last_access" json:"last_access"`
	MedecinTraitant primitive.ObjectID `bson:"doctor" json:"doctor"` //Medecin in charge of the patient
	Entity          primitive.ObjectID `bson:"entity" json:"entity"` //In which entity the DMP is created
}

// Examen : DMP Content
type Examen struct {
	ID            primitive.ObjectID   `bson:"_id" json:"-"`
	PatientRecord primitive.ObjectID   `bson:"dossier" json:"patient_record"`
	Agent         primitive.ObjectID   `bson:"agent" json:"agent"`
	Entity        primitive.ObjectID   `bson:"entity" json:"entity"`
	Content       []ExamenContent      `bson:"content" json:"content"`
	Files         []ExamenContentFiles `bson:"files" json:"files"`
	Statut        string               `bson:"statut" json:"statut"`
	Type          string               `bson:"examen_type" json:"examen_type"`
	CreatedAt     time.Time            `bson:"created_at" json:"created_at"`
}

// Antecedent : Antecedent
type Antecedent struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	PatientRecord         primitive.ObjectID `bson:"dossier" json:"patient_record"`
	Agent                 primitive.ObjectID `bson:"agent" json:"agent"`
	Entity                primitive.ObjectID `bson:"entity" json:"entity"`
	AntecedentMedical     string             `bson:"antecedent_medical" json:"antecedent_medical"`
	AntecedentChirurgical string             `bson:"antecedent_chirurgical" json:"antecedent_chirurgical"`
	AntecedentFamilial    string             `bson:"antecedent_familial" json:"antecedent_familial"`
	ModeDeVie             string             `bson:"mode_de_vie" json:"mode_de_vide"`
	CreatedAt             time.Time          `bson:"created_at" json:"created_at"`
}

// Consultation : all consultations model
type Consultation struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	PatientRecord     primitive.ObjectID `bson:"dossier" json:"patient_record"`
	Agent             primitive.ObjectID `bson:"agent" json:"agent"`
	Entity            primitive.ObjectID `bson:"entity" json:"entity"`
	MotifConsultation string             `bson:"motif_consultation" json:"motif_consultation"`
	HistoireMaladie   string             `bson:"histoire_maladie" json:"histoire_maladie"`
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
}

// Hospitalisation : all Hospitalisation model
type Hospitalisation struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	PatientRecord        primitive.ObjectID `bson:"dossier" json:"patient_record"`
	Agent                primitive.ObjectID `bson:"agent" json:"agent"`
	Entity               primitive.ObjectID `bson:"entity" json:"entity"`
	MotifHospitalisation string             `bson:"motif_hospitalisation" json:"motif_hospitalisation"`
	Commentaire          string             `bson:"commentaire" json:"commentaire"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at"`
}

// ExamenContent : Examen content
type ExamenContent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name      string             `bson:"name" json:"name"`
	Value     string             `bson:"value" json:"value"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// ExamenContentFiles : Examen content files
type ExamenContentFiles struct {
	ID        string    `bson:"_id" json:"-"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
