package route

import (
	"dna-go-app/handler"

	"github.com/gofiber/fiber/v2"

	// "dna-go-app/middleware"
	"dna-go-app/utils"
)

func RouteInit(r *fiber.App) {
	r.Get("/api/disease", handler.HandlerGetAllDisease)
	r.Post("/api/disease", utils.HandleFileUploadDisease, handler.HandlerDiseaseCreate)
	r.Get("/api/disease/:id", handler.HandlerGetDiseasebyID)
	r.Get("/api/prediction/:created_at", handler.HandlerGetResultByDate)                          //smua
	r.Get("/api/prediction/:created_at/:id_disease", handler.HandlerGetResultByDateAndIDPenyakit) //date
	r.Get("/api/prediction-last", handler.HandlerGetLastResult)
	r.Post("/api/patient/bm", utils.HandleFileUploadPatient, handler.HandlerTestDNABM)
	r.Post("/api/patient/kmp", utils.HandleFileUploadPatient, handler.HandlerTestDNAKMP)
}
