package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/platform/database"
	"strconv"
)

func GetProvinces(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}

	resp, err := db.GetProvinces()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "province not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetRegencies(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}

	provinceId, _ := strconv.Atoi(c.Params("province_id"))
	resp, err := db.GetRegencies(provinceId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "regency not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetDistricts(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}

	provinceId, _ := strconv.Atoi(c.Params("province_id"))
	regencyId, _ := strconv.Atoi(c.Params("regency_id"))
	resp, err := db.GetDistricts(provinceId, regencyId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "district not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetVillages(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}

	provinceId, _ := strconv.Atoi(c.Params("province_id"))
	regencyId, _ := strconv.Atoi(c.Params("regency_id"))
	districtId, _ := strconv.Atoi(c.Params("district_id"))
	resp, err := db.GetVillages(provinceId, regencyId, districtId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "village not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

//func GetRegency(c *fiber.Ctx) error {
//	reqs := models.RegencyRequest{}
//	resp := models.RegencyResponse{}
//	err := c.QueryParser(&reqs)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//		})
//	}
//
//	validate := utils.NewValidator()
//	if err := validate.Struct(&reqs); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//		})
//	}
//
//	db, err := database.Connect()
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//			"note":    "cant connect to database",
//		})
//	}
//
//	result, err := db.GetRegency(reqs., reqs.ProvinceId)
//
//	if err != nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error":   true,
//			"message": "regency not found",
//		})
//	}
//
//	resp = models.RegencyResponse{
//		Id:   result.Id,
//		Name: result.Name,
//	}
//	return c.Status(fiber.StatusOK).JSON(resp)
//}
//
//
//func GetVillage(c *fiber.Ctx) error {
//	reqs := models.VillageRequest{}
//	err := c.QueryParser(&reqs)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//		})
//	}
//
//	validate := utils.NewValidator()
//	if err := validate.Struct(&reqs); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//		})
//	}
//
//	db, err := database.Connect()
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error":   true,
//			"message": err.Error(),
//			"note":    "cant connect to database",
//		})
//	}
//
//	resp, err := db.GetVillages(reqs.ProvinceId, reqs.RegencyId, reqs.DistrictId)
//
//	if err != nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error":   true,
//			"message": "village not found",
//		})
//	}
//
//	return c.Status(fiber.StatusOK).JSON(resp)
//}
