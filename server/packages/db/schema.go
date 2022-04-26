package db

// import "gorm.io/gorm"

type Disease struct {
	ID          int    `json:"id"` //TODO gorm:"primarykey;autoIncrement"
	DiseaseName string `json:"disease_name"`
	DNASequence string `json:"dna_sequence"`
}

type PredictionResult struct {
	ID int `json:"id"`
	// PredictionDate 		date
	PatientName      string `json:"patient_name"`
	IDDisease        int    `json:"id_disease"`
	PredictionStatus bool   `json:"prediction_status"`
}
