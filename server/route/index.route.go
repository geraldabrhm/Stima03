package route

import (
	"github.com/gofiber/fiber/v2"
	"practice-api/handler"
	"practice-api/middleware"
)

func RouteInit(r *fiber.App) {
	r.Get("/api", middleware.Middleware, handler.HandlerGetAllDisease)
	r.Post("/api", handler.HandlerDiseaseCreate)
	r.Get("/api/:id", handler.HandlerGetDiseasebyID)
}