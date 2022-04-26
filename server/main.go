package main

import (
	"practice-api/route"
	"practice-api/database"
	"practice-api/migration"
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