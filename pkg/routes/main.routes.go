package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/pkg/middleware"
)

func MainRoutes(a *fiber.App) {
	public := a.Group("/public")
	private := a.Group("/private")
	private = private.Use(middleware.JWTProtected)
	AnnouncementRoutes(private, public)
	AuthRoutes(private, public)
	DakRoutes(public)
}
