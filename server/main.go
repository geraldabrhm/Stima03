package main

import (
	"dna-go-app/database"
	"dna-go-app/migration"
	"dna-go-app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	// Init Route
	route.RouteInit(app)
	app.Use(cors.New(cors.Config {
		AllowOrigins: "*",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))
	app.Listen(":8080")
}