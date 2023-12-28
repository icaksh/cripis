package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MainRoutes(a *fiber.App) {
	public := a.Group("/public")
	private := a.Group("/private")
	AnnouncementRoutes(private, public)
	AuthRoutes(private, public)
	LogRoutes(private, public)
	DakRoutes(public)
	TrademarkRoutes(private, public)
	UserRoutes(private, public)
}
