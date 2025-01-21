package routers

import (
	authRouters "backend/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	api := router.Group("/api")
	
	auth := api.Group("/")
	auth.Route("/", authRouters.Router)
}
