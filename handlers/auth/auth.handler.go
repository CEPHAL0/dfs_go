package authHandler

import (
	"backend/schemas"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var loginSchema schemas.LoginSchema

	err := c.BodyParser(&loginSchema)
	if err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Failed", fiber.StatusUnprocessableEntity, c)
	}

	err = utils.Validate(loginSchema)

	if err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Failed", fiber.StatusBadRequest, c)
	}

	return nil
}
