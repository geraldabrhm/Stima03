package handler

import (
	// "log"
	"dna-go-app/database"
	"dna-go-app/model/entity"
	"dna-go-app/model/request"
	"dna-go-app/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	// "fmt"
)

func HandlerGetAllDisease(ctx *fiber.Ctx) error {
	var disease []entity.Disease

	result := database.DB.Find(&disease)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed get all diseases",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successed get all diseases",
		"data":    disease,
	})
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

func HandlerTestDNABM(ctx *fiber.Ctx) error {
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

	newPatient := entity.PredictionResult{
		PatientName:      testDNA.PatientName,
		IDDisease:        int(disease.ID),
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

func HandlerTestDNAKMP(ctx *fiber.Ctx) error {
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

	newPatient := entity.PredictionResult{
		PatientName:      testDNA.PatientName,
		IDDisease:        int(disease.ID),
		PredictionStatus: true,
	}

	if utils.KMP(disease.DNASequence, fileContentStr) == -1 {
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

	return ctx.JSON(fiber.Map{
		"message": "Success get disease by id",
		"data":    disease,
	})
}

func HandlerGetResultByDate(ctx *fiber.Ctx) error {
	date := ctx.Params("created_at")
	// dateReal := date[0:10]

	var result []entity.PredictionResult

	err := database.DB.Where("date(created_at) = ?", date).Find(&result).Error
	fmt.Println(result)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get result by date",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success get result by date",
		"data":    result,
	})
}

func HandlerGetResultByDateAndIDPenyakit(ctx *fiber.Ctx) error {
	date := ctx.Params("created_at")
	diseaseId := ctx.Params("id_disease")

	var result []entity.PredictionResult

	err := database.DB.Find(&result, "date(created_at) = ? AND id_disease = ?", date, diseaseId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get result by date and disease id",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success get result by date and disease id",
		"data":    result,
	})
}

func HandlerGetLastResult(ctx *fiber.Ctx) error {
	var pred entity.PredictionResult

	err := database.DB.Last(&pred).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Failed to get last result",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success get last result",
		"data":    pred,
	})
}