package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
)

func UserProfileRoutes(public fiber.Router, private fiber.Router) {
	pathName := "/profile"
	private.Post(pathName+"/profile", controllers.RegisterProfile)
	private.Get(pathName+"/info", controllers.GetProfile)
}
