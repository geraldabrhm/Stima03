package migration

import (
	"practice-api/database"
	"practice-api/model/entity"
	"log"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Disease{}, &entity.PredictionResult{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}