package authHandler

import (
	database "backend/config"
	"backend/enums"
	authModels "backend/models/auth"
	"backend/repositories"
	"backend/schemas"
	"backend/services"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var data schemas.RegisterSchema

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(err.Error(), "Invalid Request", fiber.StatusUnprocessableEntity, c)
	}

	if err := utils.Validate(data); err != nil {
		return utils.ErrorResponse(err.Error(), "Validation Error", fiber.StatusBadRequest, c)
	}

	userRepository := repositories.NewUserRepository()
	sessionRepository := repositories.NewSessionRepository()
	userService := services.NewUserService(userRepository, sessionRepository)

	var (
		user    *authModels.User
		session *authModels.Session
	)

	// Setup Transaction to roll back if session could not be set
	err := database.WithTransaction(func(tx *gorm.DB) error {
		// Register the user and create the session as part of the transaction
		var err error
		user, session, err = userService.RegisterUser(data.Username, data.Email, data.Password, enums.CUSTOMER)
		if err != nil {
			return err
		}

		// Set the session in the user's browser
		err = utils.SetSession(tx, session, c)
		if err != nil {
			return err
		}

		// Return nil to commit transaction
		return nil
	})

	if err != nil {
		return utils.ErrorResponse(err.Error(), "Failed to Create User", fiber.StatusInternalServerError, c)
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
