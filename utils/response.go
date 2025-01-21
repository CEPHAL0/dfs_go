package utils

import (
	"backend/enums"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	StatusCode int                  `json:"statusCode"`
	Success    enums.ResponseStatus `json:"success"`
	Message    string               `json:"message"`
	Data       interface{}          `json:"data"`
}

func SuccessResponse(message string, data interface{}, status int, c *fiber.Ctx) error {
	return c.Status(status).JSON(
		APIResponse{
			StatusCode: status,
			Success:    enums.SUCCESS,
			Message:    message,
			Data:       data,
		})
}
