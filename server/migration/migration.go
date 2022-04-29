package migration

import (
	"dna-go-app/database"
	"dna-go-app/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Disease{}, &entity.PredictionResult{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}