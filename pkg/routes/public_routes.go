package routes

import (
	"github.com/kasfil/offspace/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes handle routes for unauthenticate user
func PublicRoutes(app *fiber.App) {
	// v1 routes
	v1 := app.Group("/api/v1")

	v1.Post("/signup", controllers.UserSignUp)
}
