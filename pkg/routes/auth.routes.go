package routes

import (
	"github.com/icaksh/cripis/app/controllers"
	recaptcha "github.com/jansvabik/fiber-recaptcha"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(public fiber.Router, private fiber.Router) {
	pathName := "/auth"
	public.Post(pathName+"/login", controllers.Login)
	public.Post(pathName+"/register", recaptcha.Middleware, controllers.Register)
	//
	private.Post(pathName+"/logout", controllers.Logout)
	private.Post("/token/refresh", controllers.RefreshToken)
}
