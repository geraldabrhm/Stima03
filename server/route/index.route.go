package route

import (
	"github.com/gofiber/fiber/v2"
	"practice-api/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/api", handler.HandlerGetAll)
	r.Post("/api", handler.HandlerDiseaseCreate)
}