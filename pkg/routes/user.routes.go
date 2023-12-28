package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
)

func UserRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/user"
	private.Get(pathName+"s", middleware.JWTProtected, controllers.GetUsers)
	private.Get(pathName+"/", middleware.JWTProtected, controllers.GetUser)
	private.Get(pathName+"/:id", middleware.JWTProtected, controllers.GetUserById)
	private.Put(pathName+"/", middleware.JWTProtected, controllers.EditUser)
	private.Put(pathName+"/password", middleware.JWTProtected, controllers.EditUserPassword)
	private.Put(pathName+"/status", middleware.JWTProtected, controllers.EditUserStatus)
	private.Delete(pathName+"/:id", middleware.JWTProtected, controllers.DeleteUser)
	//
	public.Get(pathName+"/roles", controllers.GetUserRoles)
}
