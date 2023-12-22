package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
)

func TrademarkRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/trademark"
	TrademarkHelpersRoutes(pathName, public)
	public.Get(pathName+"s", controllers.GetTrademarks)
	//
	private.Post(pathName+"/", controllers.CreateAnnouncement)
	private.Put(pathName+"/", controllers.UpdateAnnouncement)
	private.Delete(pathName+"/", controllers.DeleteAnnouncement)

}

func TrademarkHelpersRoutes(first string, public fiber.Router) {
	pathName := first + "/class"

	public.Get(pathName+"/", controllers.GetClasses)
	public.Get(pathName+"/:id", controllers.GetClass)
}
