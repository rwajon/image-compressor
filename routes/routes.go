package routes

import "github.com/gofiber/fiber/v2"

func Routes(router fiber.Router) {
	filesRoutes(router.Group("/files"))
}
