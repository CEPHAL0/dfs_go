package utils

import (
	enums "backend/enums"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Error struct {
	Status  enums.ResponseStatus `json:"status"`
	Code    int                  `json:"code"`
	Message string               `json:"message"`
}

func ErrorResponse(devError string, simpleError string, status int, c *fiber.Ctx) error {
	godotenv.Load()
	DEV := os.Getenv("DEV")
	if DEV == "True" {
		simpleError = devError
	}
	return c.Status(status).JSON(Error{Status: enums.FAILED, Code: status, Message: simpleError})
}
