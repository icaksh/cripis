package controllers

import (
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TrademarkSearch(c *fiber.Ctx) error {
	searchQueries := models.TrademarkFromDJKI{}
	err := c.QueryParser(&searchQueries)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&searchQueries); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	result, err := utils.FetchDataFromApi(searchQueries.Keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cannot fetch api",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result["hits"])
}

func TrademarkSimilarity(c *fiber.Ctx) error {
	searchQueries := models.TrademarkFromDJKI{}
	err := c.QueryParser(&searchQueries)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&searchQueries); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	results, err := utils.FetchDataFromApi(searchQueries.Keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cannot fetch api",
		})
	}

	hits, ok := results["hits"].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Tidak dapat mengakses data",
			"note":    "results seems invalid err: hits",
		})
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Tidak dapat mengakses data",
			"note":    "results seems invalid err: hits.hits",
		})
	}

	var brandNames []map[string]any
	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}
		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}
		brandName, ok := source["nama_merek"].(string)

		if !ok {
			continue
		}
		tClassArr, ok := source["t_class"].([]interface{})
		if !ok || len(tClassArr) == 0 {
			continue
		}

		firstTClass, ok := tClassArr[0].(map[string]interface{})
		if !ok {
			continue
		}

		classNo, ok := firstTClass["class_no"].(string)
		if !ok {
			continue
		}

		similarity := strutil.Similarity(strings.ToUpper(searchQueries.Keyword), brandName, metrics.NewHamming())

		brandNames = append(brandNames, map[string]any{"name": brandName, "class_code": classNo, "similarity": similarity})
	}
	return c.Status(fiber.StatusOK).JSON(brandNames)
}
