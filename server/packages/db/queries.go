package db

const (
	GetAllDiseases = `SELECT disease_name FROM disease`
	GetRowPenyakitID = `SELECT * FROM disease where id_disease = $X`
	GetRowPenyakitName = `SELECT * FROM disease where disease_name = $X`
	GetDNASeqDisease = `SELECT dna_sequence FROM disease where disease_name = $X`
	GetIDDisease = `SELECT id_disease FROM dna_sequence WHERE disease_name = $X`

	GetResultDate = `SELECT * FROM prediction_result WHERE id_prediction = $X`
	GetResultDateDisease =`SELECT * FROM prediction_result WHERE id_prediction = $X AND id_disease = $X`
)