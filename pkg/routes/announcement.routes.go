package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
	"github.com/icaksh/cripis/pkg/middleware"
)

func AnnouncementRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/announcement"
	public.Get(pathName+"s", controllers.GetAnnouncements)
	public.Get(pathName+"s/admin", controllers.GetAllAnnouncements)
	public.Get(pathName+"/:id", controllers.GetAnnouncement)
	//
	private.Post(pathName+"/", middleware.JWTProtected, controllers.CreateAnnouncement)
	private.Put(pathName+"/", middleware.JWTProtected, controllers.UpdateAnnouncement)
	private.Delete(pathName+"/:id", middleware.JWTProtected, controllers.DeleteAnnouncement)
}
