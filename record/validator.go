package record

// NewAntecedentPayloadValidator : for validation of antecedent creation API
type NewAntecedentPayloadValidator struct {
	AntecedentMedical     string `json:"antecedent_medical" binding:"required"`
	AntecedentChirurgical string `json:"antecedent_chirurgical" binding:"required"`
	AntecedentFamilial    string `json:"antecedent_familial" binding:"required"`
	ModeDeVie             string `json:"mode_de_vie" binding:"required"`
}

// NewConsultationPayloadValidator : for validation of antecedent creation API
type NewConsultationPayloadValidator struct {
	MotifConsultation string `json:"motif_consultation" binding:"required"`
	HistoireMaladie   string `json:"histoire_maladie" binding:"required"`
}

// NewHostpitalisationPayloadValidator : for validation of antecedent creation API
type NewHostpitalisationPayloadValidator struct {
	MotifHospitalisation string `json:"motif_hospitalisation" binding:"required"`
}

// ExamenContentValidator : for validation of examen creation API
type ExamenContentValidator struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// NewExamenValidator : for validation of examen creation API
type NewExamenValidator struct {
	Content []ExamenContent `bson:"content"  binding:"required"`
	Type    string          `json:"examen_type" binding:"required"`
}

// UpdateExamenValidator : for validation of examen udpate API
type UpdateExamenValidator struct {
	Content []ExamenContent `bson:"content" binding:"required"`
}
