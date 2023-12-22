package routes

import (
	"github.com/icaksh/cripis/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func DakRoutes(public fiber.Router) {
	public.Get("/dak/", controllers.GetProvinces)
	public.Get("/dak/:province_id", controllers.GetRegencies)
	public.Get("/dak/:province_id/:regency_id", controllers.GetDistricts)
	public.Get("/dak/:province_id/:regency_id/:district_id", controllers.GetVillages)
}
