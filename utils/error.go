package utils

import (
	enums "backend/enums"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Error struct {
	StatusCode int                  `json:"statusCode"`
	Success    enums.ResponseStatus `json:"success"`
	Message    string               `json:"message"`
}

func ErrorResponse(devError string, simpleError string, status int, c *fiber.Ctx) error {
	godotenv.Load()
	DEV := os.Getenv("DEV")
	if DEV == "True" {
		simpleError = devError
	}
	return c.Status(status).JSON(Error{StatusCode: status, Success: enums.FAILED, Message: simpleError})
}
