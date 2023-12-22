package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"time"
)

func GetClasses(c *fiber.Ctx) error {
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

func RegisterTrademark(c *fiber.Ctx) error {
	body := models.TrademarkRegistration{}
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

	Id := uuid.New()

	tmRegistration := &models.TrademarkRegistration{
		ID:                 Id,
		RegisterId:         Id,
		CreatedAt:          time.Now(),
		RegistrationNumber: "12312",
		SMECertificate:     "",
		RegisterSignature:  "",
		Status:             1,
	}
}

//
//import (
//	"github.com/gofiber/fiber/v2"
//	"github.com/icaksh/cripis/app/models"
//	"github.com/icaksh/cripis/app/utils"
//	"github.com/icaksh/cripis/platform/database"
//	"time"
//)
//
//
//func IP_Create(c *fiber.Ctx) error{
//	creds := models.SignUp{}
//	users := models.User{}
//	profile := models.UserProfile{}
//	err := c.BodyParser(&creds)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//		})
//	}
//
//	validate := utils.NewValidator()
//	if err := validate.Struct(&creds); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//		})
//	}
//
//
//	db, err := database.Connect()
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//			"note": "cant connect to database",
//		})
//	}
//
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//			"note": "cant generating salt password",
//		})
//	}
//
//	isUsernameExist := db.CheckDuplicate("username", creds.Username)
//
//	if isUsernameExist {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"error": true,
//			"note": "username has been exist",
//		})
//	}
//
//	isEmailUsed := db.CheckDuplicate("email", creds.Email)
//
//	if isEmailUsed {
//		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
//			"error": true,
//			"note": "email has been used",
//		})
//	}
//
//	profile = models.UserProfile{
//		ID: userId,
//		FirstName: creds.FirstName,
//		LastName: creds.LastName,
//	}
//
//	users = models.User{
//		ID : userId,
//		Email: creds.Email,
//		Username : creds.Username,
//		Password : string(hashedPassword),
//		CreatedAt : time.Now(),
//		Level : 0,
//		Status : 1,
//	}
//
//	if err := db.CreateUser(&users); err !=nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//			"note": "cannot store user to database",
//		})
//	}
//
//	if err := db.CreateUserProfile(&profile); err !=nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": true,
//			"message": err.Error(),
//			"note": "cannot store profile to database",
//		})
//	}
//
//	return c.Status(fiber.StatusCreated).JSON("coek")
//}
