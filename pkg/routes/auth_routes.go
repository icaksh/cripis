package routes

import (
	"github.com/icaksh/cripis/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
    // Create routes group.
	publicRoutes := a.Group("/api/v1")
	publicRoutes.Get("/search", controllers.Search)

    authRoutes := publicRoutes.Group("/auth")
	
    authRoutes.Post("/register", controllers.Register)
	authRoutes.Post("/login", controllers.Login)
	
}