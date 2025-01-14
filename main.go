package main

import (
	"backend/utils/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.SetupDatabase()
	database.Migrate(database.DB)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":3000")
}
