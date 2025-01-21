package authHandler

import (
	"backend/config"
	"backend/enums"
	"backend/repositories"
	"backend/schemas"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data schemas.RegisterSchema

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(err.Error(), "Invalid Request", fiber.StatusUnprocessableEntity, c)
	}

	if err := utils.Validate(data); err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Error", fiber.StatusBadRequest, c)
	}

	userRepository := repositories.NewUserRepository(config.DB)

	user, err := userRepository.Insert(data.Username, data.Email, data.Password, enums.CUSTOMER)

	if err != nil {
		return utils.ErrorResponse(err.Error(), "Failed to Create User", fiber.StatusBadRequest, c)
	}

	sessionRepository := repositories.NewSessionRepository(config.DB)

	newSession, err := sessionRepository.Create(user.ID)

	if err != nil {
		return utils.ErrorResponse(err.Error(), "Failed to create session", fiber.StatusInternalServerError, c)
	}

	err = sessionRepository.SetSession(newSession, c)

	if err != nil {
		return utils.ErrorResponse(err.Error(), "Failed to save session", fiber.StatusInternalServerError, c)
	}

	return utils.SuccessResponse("User Successfully Created", user, fiber.StatusCreated, c)

}

func Login(c *fiber.Ctx) error {
	var loginSchema schemas.LoginSchema

	if err := c.BodyParser(&loginSchema); err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Failed", fiber.StatusUnprocessableEntity, c)
	}

	if err := utils.Validate(loginSchema); err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Failed", fiber.StatusBadRequest, c)
	}

	return c.JSON(loginSchema)
}
