package entity

import(
	"time"
)

type Disease struct {
	ID          uint	`gorm:"primary_key;autoIncrement" json:"id"`
	DiseaseName string	`gorm:"not null;unique" json:"disease_name"`
	DNASequence string	`gorm:"not null" json:"dna_sequence"`
}

type PredictionResult struct {
	ID 					uint 		`gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt			time.Time	`gorm:"not null" json:"created_at" sql:"DEFAULT:current_timestamp"`
	PatientName      	string 		`gorm:"not null" json:"patient_name"`
	IDDisease        	int    		`gorm:"not null" json:"id_disease"`
	Disease				Disease		`gorm:"foreignKey:IDDisease"`
	PredictionStatus	bool   		`gorm:"not null" json:"prediction_status"`
}