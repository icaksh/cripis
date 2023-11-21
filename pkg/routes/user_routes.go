package routes

import (
	"github.com/icaksh/cripis/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
    // Create routes group.
    tokenRoutes := a.Group("/api/v1/token")

    // Routes for POST method:
    tokenRoutes.Post("/refresh", controllers.RefreshToken) // create a new book
}