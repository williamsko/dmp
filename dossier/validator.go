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
	AntecedentMedical     string                            `bson:"antecedent_medical" binding:"required"`
	AntecedentChirurgical string                            `bson:"antecedent_chirurgical" binding:"required"`
	AntecedentFamilial    string                            `bson:"antecedent_familial" binding:"required"`
	ModeDeVie             string                            `bson:"mode_de_vie" binding:"required"`
	Usager                usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent                 entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// NewConsultationPayloadValidator : for validation of antecedent creation API
type NewConsultationPayloadValidator struct {
	MotifConsultation string                            `bson:"motif_consultation" binding:"required"`
	HistoireMaladie   string                            `bson:"histoire_maladie" binding:"required"`
	Commentaire       string                            `bson:"commentaire" binding:"required"`
	Usager            usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent             entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// NewHostpitalisationPayloadValidator : for validation of antecedent creation API
type NewHostpitalisationPayloadValidator struct {
	MotifHospitalisation string                            `bson:"motif_hospitalisation" binding:"required"`
	Commentaire          string                            `bson:"commentaire" binding:"required"`
	Usager               usager.FindUsagerPayloadValidator `bson:"usager,omitempty" binding:"required"`
	Agent                entity.AgentPayloadValidator      `bson:"agent,omitempty" binding:"required"`
}

// ExamenContentValidator : for validation of examen creation API
type ExamenContentValidator struct {
	Name  string `bson:"name" binding:"required"`
	Value string `bson:"value" binding:"required"`
}

// NewExamenValidator : for validation of examen creation API
type NewExamenValidator struct {
	Content []ExamenContent                   `bson:"content"  binding:"required"`
	Type    string                            `bson:"examen_type" binding:"required"`
	Usager  usager.FindUsagerPayloadValidator `bson:"usager" binding:"required"`
	Agent   entity.AgentPayloadValidator      `bson:"agent" binding:"required"`
}

// UpdateExamenValidator : for validation of examen udpate API
type UpdateExamenValidator struct {
	Content []ExamenContent `bson:"content,omitempty" binding:"required"`
}
