package request

type DiseaseCreateRequest struct {
	DiseaseName string `json:"disease_name" validate:"required"`
	DNASequence string `json:"dna_sequence"`
}

type TestDNACreateRequest struct {
	PatientName      string `json:"patient_name" validate:"required"`
	IDDisease        int    `json:"id_disease" validate:"required"`
	PatientDNA       string `json:"patient_dna"`
	PredictionStatus bool   `json:"prediction_status"`
}
