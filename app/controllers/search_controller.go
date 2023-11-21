package controllers

import (
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"

	"github.com/gofiber/fiber/v2"
)



func Search(c *fiber.Ctx) error {
	searchQueries := models.SearchPDKI{}
	err := c.QueryParser(&searchQueries)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	} 

	validate := utils.NewValidator()
	if err := validate.Struct(&searchQueries); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	} 

	result, err := utils.FetchDataFromApi(searchQueries.Type, searchQueries.Keyword, searchQueries.OrderState, searchQueries.Page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
			"note": "cannot fetch api",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(result["hits"])
}