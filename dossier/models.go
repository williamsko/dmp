package dossier

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	ID             primitive.ObjectID   `json:"_id,omitempty"`
	DossierMedical primitive.ObjectID   `json:"dossier,omitempty"`
	Agent          primitive.ObjectID   `json:"agent,omitempty"`
	Entity         primitive.ObjectID   `json:"entity,omitempty"`
	Content        []ExamenContent      `json:"content,omitempty"`
	Files          []ExamenContentFiles `json:"files,omitempty"`
	Statut         string               `json:"statut,omitempty"`
	Type           string               `json:"examen_type,omitempty"`
	CreatedAt      time.Time            `json:"created_at"`
}

// Antecedent : Antecedent
type Antecedent struct {
	ID                    primitive.ObjectID `json:"_id"`
	DossierMedical        primitive.ObjectID `json:"dossier"`
	Agent                 primitive.ObjectID `json:"agent"`
	Entity                primitive.ObjectID `json:"entity"`
	AntecedentMedical     string             `json:"antecedent_medical"`
	AntecedentChirurgical string             `json:"antecedent_chirurgical"`
	AntecedentFamilial    string             `json:"antecedent_familial"`
	ModeDeVie             string             `json:"mode_de_vie"`
	CreatedAt             time.Time          `json:"created_at"`
}

// Consultation : all consultations model
type Consultation struct {
	ID                primitive.ObjectID `json:"_id"`
	DossierMedical    primitive.ObjectID `json:"dossier"`
	Agent             primitive.ObjectID `json:"agent"`
	Entity            primitive.ObjectID `json:"entity"`
	MotifConsultation string             `json:"motif_consultation"`
	HistoireMaladie   string             `json:"histoire_maladie"`
	Commentaire       string             `json:"commentaire"`
	CreatedAt         time.Time          `json:"created_at"`
}

// Hospitalisation : all Hospitalisation model
type Hospitalisation struct {
	ID                   primitive.ObjectID `json:"_id"`
	DossierMedical       primitive.ObjectID `json:"dossier"`
	Agent                primitive.ObjectID `json:"agent"`
	Entity               primitive.ObjectID `json:"entity"`
	MotifHospitalisation string             `json:"motif_hospitalisation"`
	Commentaire          string             `json:"commentaire"`
	CreatedAt            time.Time          `json:"created_at"`
}

// ExamenContent : Examen content
type ExamenContent struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Value     string             `json:"value"`
	CreatedAt time.Time          `json:"created_at"`
}

// ExamenContentFiles : Examen content files
type ExamenContentFiles struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"created_at"`
}
