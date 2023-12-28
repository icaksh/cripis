package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/platform/database"
	"strconv"
)

func GetLogs(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	result, err := db.GetLogs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get announcements, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetLogsTrademarksbyYears(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	year, _ := strconv.Atoi(c.Params("year"))

	result, err := db.GetLogsTrademarksByYears(year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get logs, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetLogsTrademarksbyMonth(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	year, _ := strconv.Atoi(c.Params("year"))
	month, _ := strconv.Atoi(c.Params("month"))
	result, err := db.GetLogsTrademarksByMonth(year, month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get logs, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetLogsUsersbyYears(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	year, _ := strconv.Atoi(c.Params("year"))

	result, err := db.GetLogsLoginByYear(year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get logs, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetLogsUsersbyMonth(c *fiber.Ctx) error {

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	year, _ := strconv.Atoi(c.Params("year"))
	month, _ := strconv.Atoi(c.Params("month"))
	result, err := db.GetLogsLoginByMonth(year, month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get logs, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
