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
			"note":    "cannot get announcements, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetAllAnnouncements(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database",
		})
	}

	result, err := db.GetAllAnnouncements()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot get announcements, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
func GetAnnouncement(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}
	id, _ := strconv.Atoi(c.Params("id"))
	result, err := db.GetAnnouncement(id)
	if err != nil {
		return utils.NotFound(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CreateAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	fmt.Println(au.AccessUuid)
	if err != nil || au.Role != 1 {
		return utils.Unauthorized(c)
	}

	body := models.AnnouncementCreation{}
	err = c.BodyParser(&body)

	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return utils.BadRequest(c, err)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	data := &models.Announcement{
		CreatedBy:   au.User,
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
	}

	err = db.CreateAnnouncement(data)

	if err != nil {
		return utils.InternalServerError(c, err)
	}

	db.CreateLog(au.User, "create announcement, id: "+c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pengumuman berhasil dibuat",
	})
}

func UpdateAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	if err != nil || au.Role != 1 {
		return utils.Unauthorized(c)
	}

	body := models.Announcement{}
	err = c.BodyParser(&body)

	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return utils.BadRequest(c, err)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	data := &models.Announcement{
		Id:          body.Id,
		Title:       body.Title,
		Description: body.Description,
		Image:       body.Image,
	}

	err = db.UpdateAnnouncement(data)

	if err != nil {
		return utils.InternalServerError(c, err)
	}

	db.CreateLog(au.User, "update announcement, id: "+c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pengumuman berhasil diubah",
	})
}

func DeleteAnnouncement(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	if err != nil || au.Role != 1 {
		return utils.Unauthorized(c)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	id, _ := strconv.Atoi(c.Params("id"))

	res := db.DeleteAnnouncement(id)
	if res != nil {
		return utils.InternalServerError(c, res)
	}

	db.CreateLog(au.User, "delete announcement, id: "+c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pengumuman berhasil dihapus",
	})
}
