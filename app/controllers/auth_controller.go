package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login method for create a new user
// @Description Create a new user
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {string} status "created"
// @Router /v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	if c.Locals("recaptchaSuccess") == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Captcha tidak dapat divalidasi",
			"note":    "invalid captcha",
		})
	}

	creds := models.Credentials{}
	err := c.BodyParser(&creds)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat diparse, mohon cek kembali",
			"note":    err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
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

	auth, err := db.Auth(creds.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Email yang Anda gunakan tidak ditemukan",
			"note":    "wrong username",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Password yang Anda masukkan salah",
			"note":    "wrong password",
		})
	}

	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TIME_KEY_EXPIRE_HOURS_COUNT"))

	if !creds.Remember {
		hoursCount, _ = strconv.Atoi(os.Getenv("JWT_REFRESH_TIME_KEY_EXPIRE_HOURS_COUNT_REMEMBER"))
	}
	refreshTime := time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()
	au := &models.JwtAuthModel{
		AccessId:  uuid.New(),
		RefreshId: uuid.New(),
		UserId:    auth.ID,
		Email:     auth.Email,
		FirstName: auth.FirstName,
		LastName:  auth.LastName,
		Role:      auth.Level,
		Duration:  refreshTime,
	}
	token, err := utils.GenerateNewAuthToken(au)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create jwt token",
		})
	}

	redisDb, err := database.RedisConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to redis",
		})
	}
	refresh := redisDb.CreateAuth(auth.ID, token)
	if refresh != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create jwt token",
		})
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	return c.Status(fiber.StatusCreated).JSON(tokens)
}

func RefreshToken(c *fiber.Ctx) error {
	body := models.RefreshToken{}
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

	redisDb, err := database.RedisConnect()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    err,
		})
	}
	claims, err := utils.ExtractRefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error) ",
			"note":    err,
		})
	}

	deleted, err := redisDb.DeleteAuth(claims.RefreshUuid)

	if err != nil || deleted == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Token tidak valid",
			"note":    err,
		})
	}

	au := &models.JwtAuthModel{
		AccessId:  uuid.New(),
		RefreshId: uuid.New(),
		UserId:    claims.User,
		Email:     claims.Email,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Role:      claims.Role,
		Duration:  claims.ExpiresAt,
	}

	token, err := utils.GenerateNewAuthToken(au)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error) ",
			"note":    err.Error(),
		})
	}

	refresh := redisDb.CreateAuth(claims.User, token)
	if refresh != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create jwt token",
		})
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	return c.Status(fiber.StatusCreated).JSON(tokens)
}

func Logout(c *fiber.Ctx) error {

	au, err := utils.ExtractTokenMetadata(c)
	fmt.Println(au.AccessUuid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}
	redisDb, err := database.RedisConnect()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"note":    "Terjadi kesalahan (Internal Server Error)",
			"message": err,
		})
	}

	deleted, err := redisDb.DeleteAuth(au.AccessUuid)
	if err != nil || deleted == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil keluar",
	})
}
