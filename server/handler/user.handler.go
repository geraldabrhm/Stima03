package handler

import (
	"github.com/gofiber/fiber/v2"
	"practice-api/database"
	"practice-api/model/entity"
	"practice-api/model/request"
	"log"
	"fmt"
)

func HandlerRead(ctx *fiber.Ctx) error {

	return ctx.Status(500).JSON(fiber.Map{
		"Hello": "world",
	})
}

func HandlerGetAll(ctx *fiber.Ctx) error {
	var disease []entity.Disease

	result := database.DB.Find(&disease)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&disease).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.Status(500).JSON(disease)
}

func HandlerDiseaseCreate(ctx *fiber.Ctx) error {
	disease := new(request.DiseaseCreateRequest)
	fmt.Println(disease)

	if err := ctx.BodyParser(disease); err != nil {
		return err
	}

	newDisease := entity.Disease {
		DiseaseName	: disease.DiseaseName,
		DNASequence	: disease.DNASequence,
	}
	fmt.Println(newDisease)

	errCreateDisease := database.DB.Create(&newDisease).Error

	if errCreateDisease != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to insert the disease",
		}) 
	}

	return ctx.JSON(fiber.Map{
		"message": "success insert the new disease",
		"data": newDisease,
	})
}