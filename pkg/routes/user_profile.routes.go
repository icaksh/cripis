package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
)

func UserProfileRoutes(public fiber.Router, private fiber.Router) {
	pathName := "/profile"
	private.Post(pathName+"/profile", middleware.JWTProtected, controllers.RegisterProfile)
	private.Get(pathName+"/info", middleware.JWTProtected, controllers.GetProfile)
}
