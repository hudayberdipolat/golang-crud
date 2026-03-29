package routes

import "github.com/gofiber/fiber/v3"

func SetAllRoutes(app *fiber.App) {
	setAuthRoutes(app)
	setUserRoutes(app)
	setPostRoutes(app)
}
