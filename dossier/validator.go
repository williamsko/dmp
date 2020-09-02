package dossier

import (
	"dmp/entity"
	"dmp/usager"
)

// NewDossierPayloadValidator : for validation of dossier creation API
type NewDossierPayloadValidator struct {
	Usager usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent  entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}

// NewAntecedentPayloadValidator : for validation of antecedent creation API
type NewAntecedentPayloadValidator struct {
	AntecedentMedical     string                            `bson:"antecedent_medical,omitempty" json:"antecedent_medical,omitempty" binding:"required"`
	AntecedentChirurgical string                            `bson:"antecedent_chirurgical,omitempty" json:"antecedent_chirurgical,omitempty" binding:"required"`
	AntecedentFamilial    string                            `bson:"antecedent_familial,omitempty" json:"antecedent_familial,omitempty"  binding:"required"`
	ModeDeVie             string                            `bson:"mode_de_vie,omitempty" json:"mode_de_vie,omitempty"  binding:"required"`
	Usager                usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent                 entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}

// NewConsultationPayloadValidator : for validation of antecedent creation API
type NewConsultationPayloadValidator struct {
	MotifConsultation string                            `bson:"motif_consultation,omitempty" json:"motif_consultation,omitempty" binding:"required"`
	HistoireMaladie   string                            `bson:"histoire_maladie,omitempty" json:"histoire_maladie,omitempty" binding:"required"`
	Commentaire       string                            `bson:"commentaire,omitempty" json:"commentaire,omitempty"  binding:"required"`
	Usager            usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent             entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}

// NewHostpitalisationPayloadValidator : for validation of antecedent creation API
type NewHostpitalisationPayloadValidator struct {
	MotifHospitalisation string                            `bson:"motif_hospitalisation,omitempty" json:"motif_hospitalisation,omitempty" binding:"required"`
	Commentaire          string                            `bson:"commentaire,omitempty" json:"commentaire,omitempty"  binding:"required"`
	Usager               usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent                entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}

// ExamenContentValidator : for validation of examen creation API
type ExamenContentValidator struct {
	Name  string `bson:"name,omitempty" json:"name,omitempty"  binding:"required"`
	Value string `bson:"value,omitempty" json:"value,omitempty"  binding:"required"`
}

// NewExamenValidator : for validation of examen creation API
type NewExamenValidator struct {
	Content []ExamenContent                   `bson:"content,omitempty" json:"content,omitempty"  binding:"required"`
	Type    string                            `bson:"examen_type,omitempty" json:"examen_type,omitempty"  binding:"required"`
	Usager  usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent   entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}
