package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
)

func GetTrademarks(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}

	queryValue := c.Query("name")
	query := []models.Trademark{}
	if len(queryValue) == 0 {
		query, err = db.GetAllTrademarks()
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Merek dagang tidak ditemukan",
				"note":    "cannot get trademark err:" + err.Error(),
			})
		}
	} else {
		query, err = db.GetAllTrademarksByName(queryValue)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Merek dagang tidak ditemukan",
				"note":    "cannot get trademark err:" + err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

func CreateTrademark(c *fiber.Ctx) error {
	_, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}

	body := models.Trademark{}
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat diparse, mohon cek kembali",
			"note":    err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat divalidasi, mohon cek kembali",
			"note":    err.Error(),
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	data := &models.Trademark{
		ID:             uuid.New(),
		Name:           body.Name,
		Class:          body.Class,
		RegistrationId: body.RegistrationId,
		Status:         1,
	}

	query := db.CreateTrademark(data)
	if query != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create trademark",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "trademark successfully created",
	})
}
