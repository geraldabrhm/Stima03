package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Middleware(ctx *fiber.Ctx) error{
	authorizeToken := ctx.Get("x-token")

	if authorizeToken != "FERGUSO" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Doesn't have authorization",
		})
	}

	return ctx.Next()
}