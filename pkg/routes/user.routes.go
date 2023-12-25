package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
)

func UserRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/user"
	public.Get(pathName+"/roles", controllers.GetUserRoles)
	//
	private.Get(pathName+"s", middleware.JWTProtected, controllers.GetUsers)
	private.Get(pathName+"/:id", middleware.JWTProtected, controllers.GetProfilebyId)
}
