package route

import (
	"practice-api/handler"

	"github.com/gofiber/fiber/v2"

	// "practice-api/middleware"
	"practice-api/utils"
)

func RouteInit(r *fiber.App) {
	r.Get("/api/disease", handler.HandlerGetAllDisease)
	r.Post("/api/disease", utils.HandleFileUploadDisease, handler.HandlerDiseaseCreate)
	r.Get("/api/disease/:id", handler.HandlerGetDiseasebyID)
	r.Get("/api/prediction/:date", handler.HandlerGetResultByDate)                          //smua
	r.Get("/api/prediction/:date/:idpenyakit", handler.HandlerGetResultByDateAndIDPenyakit) //date
}
