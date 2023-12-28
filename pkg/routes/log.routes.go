package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/controllers"
)

func LogRoutes(private fiber.Router, public fiber.Router) {
	pathName := "/log"
	private.Get(pathName+"s", controllers.GetLogs)
	private.Get(pathName+"/:year/:month", controllers.GetLogsTrademarksbyMonth)
	//
	public.Get(pathName+"/trademark/:year", controllers.GetLogsTrademarksbyYears)
	public.Get(pathName+"/trademark/:year/:month", controllers.GetLogsTrademarksbyMonth)
	public.Get(pathName+"/login/:year", controllers.GetLogsUsersbyYears)
	public.Get(pathName+"/login/:year/:month", controllers.GetLogsUsersbyMonth)
}
