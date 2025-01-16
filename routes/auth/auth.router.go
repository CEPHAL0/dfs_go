package authRouter

import (
	authHandler "backend/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/login", authHandler.Login)
}
