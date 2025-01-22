package main

import (
	database "backend/config"
	routers "backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.SetupDatabase()
	database.Initialize(database.DB)
	database.Migrate(database.DB)

	routers.Initialize(app)

	app.Listen(":3000")
}
