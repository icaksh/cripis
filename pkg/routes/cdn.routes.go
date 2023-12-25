package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/pkg/middleware"
)

func CdnRoutes(private fiber.Router) {
	pathName := "/cdn"
	private.Post(pathName+"/logout", middleware.AccessControl())
}
