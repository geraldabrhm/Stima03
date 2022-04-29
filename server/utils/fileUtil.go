package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"fmt"
	"os"
	"bufio"
)

func HandleFileUploadDisease(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("DNASequence")

	if errFile != nil {
		log.Println("Error handle file disease input", errFile)
	}

	var filename *string

	if file != nil {
		fileType := file.Header.Get("Content-Type")

		if fileType != "text/plain" {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Doesn't support file upload rather than .txt file",
			})
		}

		filename = &file.Filename
		
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/disease/%s", *filename))

		if errSaveFile != nil {
			log.Println("Error handle file disease saving")
		}
	} else {
		log.Println("No file to upload")
	}

	if filename != nil {
		fileCt, errCt := os.Open(fmt.Sprintf("./public/disease/%s", *filename))

		if errCt != nil {
			panic(errCt)
		}

		scanner := bufio.NewScanner(fileCt)

		var filecontent string

		for scanner.Scan() {
			filecontent = scanner.Text()
		}

		if filecontent == "" {
			ctx.Locals("filecontent", nil)
		}
		ctx.Locals("filecontent", filecontent)
	} else {
		ctx.Locals("filecontent", nil)
	}
	
	return ctx.Next()
}

func HandleFileUploadPatient(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("PatientDNA")

	if errFile != nil {
		log.Println("Error handle file patient DNA input", errFile)
	}

	var filename *string

	if file != nil {
		fileType := file.Header.Get("Content-Type")

		if fileType != "text/plain" {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Doesn't support file upload rather than .txt file",
			})
		}

		filename = &file.Filename
		
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/patient_dna/%s", *filename))

		if errSaveFile != nil {
			log.Println("Error handle file patient DNA saving")
		}
	} else {
		log.Println("No file to upload")
	}

	if filename != nil {
		fileCt, errCt := os.Open(fmt.Sprintf("./public/patient_dna/%s", *filename))

		if errCt != nil {
			panic(errCt)
		}

		scanner := bufio.NewScanner(fileCt)

		var filecontent2 string

		for scanner.Scan() {
			filecontent2 = scanner.Text()
		}

		if filecontent2 == "" {
			ctx.Locals("filecontent2", nil)
		}
		ctx.Locals("filecontent2", filecontent2)
	} else {
		ctx.Locals("filecontent2", nil)
	}
	
	return ctx.Next()
}