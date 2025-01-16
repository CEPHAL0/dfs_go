package main

import (
	database "backend/config"
	router "backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.SetupDatabase()
	// database.Migrate(database.DB)

	router.Initialize(app)

	app.Listen(":3000")
}
