package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
)

func AnnouncementRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/announcements"
	public.Get(pathName+"/", controllers.GetAnnouncements)
	public.Get(pathName+"/:id", controllers.GetAnnouncement)
	//
	private.Post(pathName+"/", controllers.CreateAnnouncement)
	private.Put(pathName+"/", controllers.UpdateAnnouncement)
	private.Delete(pathName+"/", controllers.DeleteAnnouncement)
}
