package routes

import (
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
	recaptcha "github.com/jansvabik/fiber-recaptcha"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/auth"
	public.Post(pathName+"/login", recaptcha.Middleware, controllers.Login)
	public.Post(pathName+"/refresh", controllers.RefreshToken)
	public.Post(pathName+"/register", recaptcha.Middleware, controllers.Register)
	public.Post(pathName+"/reset", recaptcha.Middleware, controllers.Reset)
	//
	private.Post(pathName+"/logout", middleware.JWTProtected, controllers.Logout)
}
