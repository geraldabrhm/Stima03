package main

import (
	"dna-go-app/database"
	"dna-go-app/migration"
	"dna-go-app/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	// Init Route
	route.RouteInit(app)

	app.Listen(":8080")
}