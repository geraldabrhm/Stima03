package handler

import (
	"log"
	"practice-api/database"
	"practice-api/model/entity"
	"practice-api/model/request"
	"practice-api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
			"error":   errValid.Error(),
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

	newDisease := entity.Disease{
		DiseaseName: disease.DiseaseName,
		DNASequence: fileContentStr,
	}

	errCreateDisease := database.DB.Create(&newDisease).Error

	if errCreateDisease != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to insert the disease, it might be because the disease name is already exist",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successed insert the new disease",
		"data":    newDisease,
	})
}

func HandlerTestDNA(ctx *fiber.Ctx) error {
	testDNA := new(request.TestDNACreateRequest)

	if err := ctx.BodyParser(testDNA); err != nil {
		return err
	}

	valid := validator.New()

	errValid := valid.Struct(testDNA)

	if errValid != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed because required file does not filled",
			"error":   errValid.Error(),
		})
	}

	var fileContentStr string
	fileContent := ctx.Locals("filecontent2")

	if fileContent == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "File DNA Sequence is required to test DNA or your file might be empty",
		})
	} else {
		fileContentStr = fileContent.(string)

		if !utils.CheckDNAInput(fileContentStr) {
			return ctx.Status(422).JSON(fiber.Map{
				"message": "DNA is unacceptable, please check again the DNA Sequence",
			})
		}

	}

	var disease entity.Disease

	err := database.DB.First(&disease, "id = ?", testDNA.IDDisease).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get disease by id",
		})
	}

	newPatient := request.TestDNACreateRequest{
		PatientName:      testDNA.PatientName,
		IDDisease:        testDNA.IDDisease,
		PatientDNA:       fileContentStr,
		PredictionStatus: true,
	}

	if utils.BoyerMoore(disease.DNASequence, fileContentStr) == -1 {
		newPatient.PredictionStatus = false
	}

	errCreatePatientValue := database.DB.Create(&newPatient).Error

	if errCreatePatientValue != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to insert new prediction value",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successed insert the new prediction",
		"data":    newPatient,
	})
}

func HandlerGetDiseasebyID(ctx *fiber.Ctx) error {
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
		"data":    disease,
	})
}

func HandlerGetResultByDate(ctx *fiber.Ctx) error {
	date := ctx.Params("created_at")
	dateReal := date[0:10]

	var result entity.PredictionResult

	err := database.DB.Find(&result, "date(created_at) = ?", dateReal).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get result by date",
		})
	}

	return ctx.Status(404).JSON(fiber.Map{
		"message": "Success get result by date",
		"data":    result,
	})
}

func HandlerGetResultByDateAndIDPenyakit(ctx *fiber.Ctx) error {
	date := ctx.Params("created_at")
	dateReal := date[0:10]
	diseaseId := ctx.Params("id_disease")

	var result entity.PredictionResult

	err := database.DB.Find(&result, "date(created_at) = ? AND id_disease = ?", dateReal, diseaseId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get result by date and disease id",
		})
	}

	return ctx.Status(404).JSON(fiber.Map{
		"message": "Success get result by date and disease id",
		"data":    result,
	})
}
