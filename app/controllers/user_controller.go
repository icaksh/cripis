package controllers

import (
	"fmt"
	"time"

	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register method for create a new user
// @Description Create a new user
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {string} status "created"
// @Router /v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	//if c.Locals("recaptchaSuccess") == false {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"error":   true,
	//		"note":    "Captcha yang Anda berikan tidak valid",
	//		"message": "invalid captcha",
	//	})
	//}
	creds := models.SignUp{}
	users := models.User{}
	err := c.BodyParser(&creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"note":    "Mohon cek kembali data yang Anda masukkan",
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"note":    "Terjadi kesalahan (Internal Server Error)",
			"message": err.Error(),
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"note":    "Terjadi kesalahan (Internal Server Error)",
			"message": "cant connect to database",
		})
	}

	userId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "cant generating salt password",
			"note":    "Terjadi kesalahan (Internal Server Error)",
		})
	}

	isEmailUsed := db.CheckDuplicateUsers("email", creds.Email)

	if isEmailUsed {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   true,
			"note":    "Email telah didaftarkan",
			"message": "email has been used",
		})
	}

	users = models.User{
		ID:        userId,
		Email:     creds.Email,
		FirstName: creds.FirstName,
		LastName:  creds.LastName,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		Level:     0,
		Status:    1,
	}

	if err := db.CreateUser(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "cannot store user to database",
			"note":    "Terjadi kesalahan (Internal Server Error)",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(users)
}

func RegisterProfile(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}
	reqs := models.UserProfile{}
	profile := models.UserProfile{}
	err = c.BodyParser(&reqs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&reqs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak valid, mohon cek kembali",
			"note":    err.Error(),
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

	isProfileExist, err := db.CheckDuplicateProfile("user_id", claims.User)

	fmt.Println(isProfileExist)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "cannot check duplicate",
			"note":    "Terjadi kesalahan (Internal Server Error)",
		})
	}

	if isProfileExist {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   true,
			"message": "Profil telah ada",
			"note":    "profile has been exist",
		})
	}

	profile = models.UserProfile{
		ID:         claims.User,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CardNumber: reqs.CardNumber,
		Address:    reqs.Address,
		DoB:        reqs.DoB,
		Sex:        reqs.Sex,
		Province:   reqs.Province,
		Regency:    reqs.Regency,
		District:   reqs.District,
		Village:    reqs.Village,
		PostalCode: reqs.PostalCode,
	}

	if err := db.CreateUserProfile(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"note":    "cannot store user to database",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(profile)
}

func GetProfile(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	db, err := database.Connect()
	profile, err := db.GetUserProfile(claims.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"note":    "Profil Akun tidak ditemukan",
			"message": "no profile",
		})
	}
	return c.Status(fiber.StatusOK).JSON(profile)

}
