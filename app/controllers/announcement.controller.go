package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"strconv"
)

func GetAnnouncements(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	result, err := db.GetAnnouncements()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get announcements",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetAnnouncement(c *fiber.Ctx) error {
	body := models.Announcement{}
	err := c.BodyParser(&body)

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

	result, err := db.GetAnnouncement(body.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get announcement from id:" + strconv.Itoa(body.Id),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CreateAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	fmt.Println(au.AccessUuid)
	if err != nil || au.Role != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}

	body := models.Announcement{}
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

	data := &models.Announcement{
		CreatedBy:   au.User,
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
	}

	result, err := db.UpdateAnnouncement(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot update announcement from id:" + strconv.Itoa(body.Id),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func UpdateAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	if err != nil || au.Role != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}

	body := models.Announcement{}
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

	data := &models.Announcement{
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
	}

	result, err := db.UpdateAnnouncement(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot update announcement from id:" + strconv.Itoa(body.Id),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	if err != nil || au.Role != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}

	body := models.Announcement{}
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

	data := &models.Announcement{
		CreatedBy:   au.User,
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
	}

	res := db.DeleteAnnouncement(id)
	if res != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pengumuman berhasil dihapus",
	})
}
