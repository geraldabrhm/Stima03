package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"practice-api/database"
	"practice-api/model/entity"
	"practice-api/model/request"
	"log"
)

func HandlerRead(ctx *fiber.Ctx) error {

	return ctx.Status(500).JSON(fiber.Map{
		"Hello": "world",
	})
}

func HandlerGetAllDisease(ctx *fiber.Ctx) error {
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

	if err := ctx.BodyParser(disease); err != nil {
		return err
	}

	valid := validator.New()

	errValid := valid.Struct(disease)

	if errValid != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed because required form does not filled",
			"error": errValid.Error(),
		})
	}

	newDisease := entity.Disease {
		DiseaseName	: disease.DiseaseName,
		DNASequence	: disease.DNASequence,
	}

	errCreateDisease := database.DB.Create(&newDisease).Error

	if errCreateDisease != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to insert the disease",
		}) 
	}

	return ctx.JSON(fiber.Map{
		"message": "Successed insert the new disease",
		"data": newDisease,
	})
}

func HandlerGetDiseasebyID(ctx* fiber.Ctx) error {
	diseaseId := ctx.Params("id")

	var disease entity.Disease

	err := database.DB.First(&disease, "id = ?", diseaseId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get disease by id",
		})
	}

	return ctx.Status(404).JSON(fiber.Map{
		"message": "Success get disease by id",
		"data": disease,
	})
}