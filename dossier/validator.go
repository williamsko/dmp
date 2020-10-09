package dossier

import (
	"dmp/entity"
	"dmp/usager"
)

// NewDossierPayloadValidator : for validation of dossier creation API
type NewDossierPayloadValidator struct {
	Usager usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent  entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// NewAntecedentPayloadValidator : for validation of antecedent creation API
type NewAntecedentPayloadValidator struct {
	AntecedentMedical     string                            `json:"antecedent_medical" binding:"required"`
	AntecedentChirurgical string                            `json:"antecedent_chirurgical" binding:"required"`
	AntecedentFamilial    string                            `json:"antecedent_familial" binding:"required"`
	ModeDeVie             string                            `json:"mode_de_vie" binding:"required"`
	Agent                 entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// NewConsultationPayloadValidator : for validation of antecedent creation API
type NewConsultationPayloadValidator struct {
	MotifConsultation string                            `json:"motif_consultation" binding:"required"`
	HistoireMaladie   string                            `json:"histoire_maladie" binding:"required"`
	Usager            usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent             entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// NewHostpitalisationPayloadValidator : for validation of antecedent creation API
type NewHostpitalisationPayloadValidator struct {
	MotifHospitalisation string                            `json:"motif_hospitalisation" binding:"required"`
	Usager               usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent                entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// ExamenContentValidator : for validation of examen creation API
type ExamenContentValidator struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// NewExamenValidator : for validation of examen creation API
type NewExamenValidator struct {
	Content []ExamenContent                   `bson:"content"  binding:"required"`
	Type    string                            `json:"examen_type" binding:"required"`
	Usager  usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent   entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// UpdateExamenValidator : for validation of examen udpate API
type UpdateExamenValidator struct {
	Content []ExamenContent `bson:"content" binding:"required"`
}
