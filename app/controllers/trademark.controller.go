package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"strings"
	"time"
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

func GetTrademarksByUser(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cant connect to database",
		})
	}
	res, err := db.GetTrademarksByUser(at.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
			"note":    "cannot get trademark err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func CreateTrademark(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
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
	currentTime := time.Now()
	twoLetterMonth := currentTime.Format("Jan")[:2]
	formattedTime := currentTime.Format("06010215040506")
	registrationNumber := strings.ToUpper(twoLetterMonth) + "00" + formattedTime
	fmt.Println(registrationNumber)
	data := &models.Trademark{
		ID:             uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      at.User,
		RegisterNumber: registrationNumber,
		TrademarkName:  body.TrademarkName,
		Class:          body.Class,
		OwnerName:      body.OwnerName,
		Address:        body.Address,
		Province:       body.Province,
		Regency:        body.Regency,
		District:       body.District,
		Village:        body.Village,
		Image:          body.Image,
		Status:         1,
	}

	query := db.CreateTrademark(data)
	if query != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create trademark, err: " + query.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "trademark successfully created",
	})
}
