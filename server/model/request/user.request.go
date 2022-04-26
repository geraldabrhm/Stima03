package request

type DiseaseCreateRequest struct {
	DiseaseName string	`json:"disease_name" validate:"required"`
	DNASequence string	`json:"dna_sequence" validate:"required"`
}

type PredictionResultCreateRequest struct {
	PatientName      	string 		`json:"patient_name" validate:"required"`
	IDDisease        	int    		`json:"id_disease" validate:"required"`
	PredictionStatus	bool   		`json:"prediction_status" validate:"required"`
}