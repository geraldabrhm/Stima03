package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"practice-api/database"
	"practice-api/model/entity"
	"practice-api/model/request"
	"practice-api/utils"
	"log"
	// "fmt"
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
			"message": "Failed because required file does not filled",
			"error": errValid.Error(),
		})
	}

	var fileContentStr string
	fileContent := ctx.Locals("filecontent")

	if fileContent == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "File is required to submit new disease or your file might be empty",
		})
	} else {
		fileContentStr = fileContent.(string)

		if !utils.CheckDNAInput(fileContentStr) {
			return ctx.Status(422).JSON(fiber.Map{
				"message": "DNA is unacceptable, please check again the DNA Sequence",
			})
		}
	}

	newDisease := entity.Disease {
		DiseaseName	: disease.DiseaseName,
		DNASequence	: fileContentStr,
	}

	errCreateDisease := database.DB.Create(&newDisease).Error

	if errCreateDisease != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to insert the disease, it might be because the disease name is already exist",
		}) 
	}

	return ctx.JSON(fiber.Map{
		"message": "Successed insert the new disease",
		"data": newDisease,
	})
}

// func HandlerTestDNA(ctx *fiber.Ctx) error {
// 	testDNA := new(request.TestDNACreateRequest)

// 	if err := ctx.BodyParser(testDNA); err != nil {
// 		return err
// 	}

// 	valid := validator.New()

// 	errValid := valid.Struct(testDNA)

// 	if errValid != nil {
// 		return ctx.Status(400).JSON(fiber.Map{
// 			"message": "Failed because required file does not filled",
// 			"error": errValid.Error(),
// 		})
// 	}

// 	var fileContentStr string
// 	fileContent := ctx.Locals("filecontent2")

// 	if fileContent == nil {
// 		return ctx.Status(422).JSON(fiber.Map{
// 			"message": "File is required to submit new disease or your file might be empty",
// 		})	
// 	} else {
// 		fileContentStr = fileContent.(string)

// 		if !utils.CheckDNAInput(fileContentStr) {
// 			return ctx.Status(422).JSON(fiber.Map{
// 				"message": "DNA is unacceptable, please check again the DNA Sequence",
// 			})
// 		}
// 	}

// 	newPatient := request.TestDNACreateRequest {
// 		PatientName	: testDNA.PatientName,
// 		PatientDNA : fileContentStr,
// 	}

// 	errCreateDisease := database.DB.Create(&newPatient).Error

// 	if errCreateDisease != nil {
// 		return ctx.Status(500).JSON(fiber.Map{
// 			"message": "Failed to insert the disease, it might be because the disease name is already exist",
// 		}) 
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"message": "Successed insert the new disease",
// 		"data": newPatient,
// 	})
// }

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