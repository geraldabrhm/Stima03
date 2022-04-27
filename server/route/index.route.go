package route

import (
	"github.com/gofiber/fiber/v2"
	"practice-api/handler"
	// "practice-api/middleware"
	"practice-api/utils"
)

func RouteInit(r *fiber.App) {
	r.Get("/api/disease", handler.HandlerGetAllDisease)
	r.Post("/api/disease", utils.HandleFileUploadDisease, handler.HandlerDiseaseCreate)
	r.Get("/api/disease/:id", handler.HandlerGetDiseasebyID)	
}