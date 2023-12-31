package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
)

func TrademarkRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/trademark"
	TrademarkHelpersRoutes(pathName, public)
	TrademarkSearch(pathName, public)
	TrademarkRegister(pathName, private)

	public.Get(pathName+"s", controllers.GetTrademarksBySearch)
	private.Get(pathName+"s/admin", middleware.JWTProtected, controllers.GetTrademarks)

	private.Get(pathName+"s", middleware.JWTProtected, controllers.GetTrademarksByUser)
	//
	private.Get(pathName+"/:id", middleware.JWTProtected, controllers.GetTrademarkById)
	private.Post(pathName+"/", middleware.JWTProtected, controllers.CreateTrademark)
	private.Put(pathName+"/", middleware.JWTProtected, controllers.UpdateTrademark)
	private.Put(pathName+"/status", middleware.JWTProtected, controllers.UpdateTrademarkStatus)
	private.Delete(pathName+"/:id", middleware.JWTProtected, controllers.DeleteTrademark)

}

func TrademarkHelpersRoutes(first string, public fiber.Router) {
	public.Get(first+"/class", controllers.GetClasses)
	public.Get(first+"/class/:id", controllers.GetClass)
	public.Get(first+"/status", controllers.GetTrademarkStatus)
}

func TrademarkSearch(first string, public fiber.Router) {
	pathName := first + "/search"
	public.Get(pathName+"/", controllers.TrademarkSimilarity)
	public.Get(pathName+"/:id", controllers.GetClass)
}

func TrademarkRegister(first string, private fiber.Router) {
	pathName := first + "/register"
	private.Post(pathName+"/", middleware.JWTProtected, controllers.TrademarkSimilarity)
}
