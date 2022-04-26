package request

// import(
// 	"gorm.io/gorm"
// 	"time"
// )

type DiseaseCreateRequest struct {
	DiseaseName string	`json:"disease_name"`
	DNASequence string	`json:"dna_sequence"`
}

type PredictionResultCreateRequest struct {
	PatientName      	string 		`json:"patient_name"`
	IDDisease        	int    		`json:"id_disease"`
	PredictionStatus	bool   		`json:"prediction_status"`
}