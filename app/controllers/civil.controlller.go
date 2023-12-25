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

func GetAddressFromDak(c *fiber.Ctx) error {
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
	villageId, _ := strconv.Atoi(c.Params("village_id"))
	resp, err := db.GetAddressFromDak(provinceId, regencyId, districtId, villageId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "address not found, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
