package entity

import(
	"time"
)

type Disease struct {
	ID          uint		`json:'id'`
	DiseaseName string	`json:"disease_name"`
	DNASequence string	`json:"dna_sequence"`
}

type PredictionResult struct {
	ID 					uint 		`json:"id"`
	CreatedAt			time.Time	`json:"created_at"`
	PatientName      	string 		`json:"patient_name"`
	IDDisease        	int    		`json:"id_disease"`
	PredictionStatus	bool   		`json:"prediction_status"`
}