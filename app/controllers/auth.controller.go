package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"github.com/thanhpk/randstr"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	if c.Locals("recaptchaSuccess") == false {
		return utils.BadRequest(c, fmt.Errorf("captcha yang Anda berikan tidak valid"))
	}

	creds := models.Credentials{}
	err := c.BodyParser(&creds)
	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
		return utils.BadRequest(c, err)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.BadRequest(c, err)
	}

	auth, err := db.Auth(creds.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Email/Password yang Anda gunakan tidak ditemukan",
			"note":    "wrong email",
		})
	}

	if auth.Verified != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Akun Anda belum terverifikasi, silakan hubungi Admin",
			"note":    "wrong email",
		})
	}

	if auth.Status != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Akun Anda dalam kondisi nonaktif, silakan hubungi Admin",
			"note":    "wrong email",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Email/Password yang Anda gunakan tidak ditemukan",
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
		Role:      auth.Roles,
		Duration:  refreshTime,
	}
	token, err := utils.GenerateNewAuthToken(au)
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	redisDb, err := database.RedisConnect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}
	refresh := redisDb.CreateAuth(auth.ID, token)
	if refresh != nil {
		return utils.InternalServerError(c, err)
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	db.CreateLog(auth.ID, "login, ip: "+c.IP()+", user agent: "+c.Get("User-Agent"))
	return c.Status(fiber.StatusCreated).JSON(tokens)
}

func Register(c *fiber.Ctx) error {
	if c.Locals("recaptchaSuccess") == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "invalid captcha",
			"message": "Captcha yang Anda berikan tidak valid",
		})
	}
	creds := models.SignUp{}
	users := models.User{}
	err := c.BodyParser(&creds)
	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
		return utils.BadRequest(c, err)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	userId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	if err != nil {
		return utils.InternalServerError(c, err)
	}

	isEmailUsed, err := db.CheckDuplicateUsers("email", creds.Email)
	if err != nil {
		return utils.InternalServerError(c, err)
	}
	if isEmailUsed {
		return utils.Conflict(c, err, "Email")
	}

	users = models.User{
		ID:         userId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Email:      creds.Email,
		FirstName:  strings.ToUpper(creds.FirstName),
		LastName:   strings.ToUpper(creds.LastName),
		Password:   string(hashedPassword),
		Roles:      2,
		Verified:   2,
		Status:     1,
		CardNumber: creds.CardNumber,
		Address:    creds.Address,
		DoB:        creds.DoB,
		Sex:        creds.Sex,
		Village:    creds.Village,
		District:   creds.District,
		Regency:    creds.Regency,
		Province:   creds.Province,
		PostalCode: creds.PostalCode,
	}

	if err := db.CreateUser(&users); err != nil {
		return utils.InternalServerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(users)
}

func Reset(c *fiber.Ctx) error {
	if c.Locals("recaptchaSuccess") == false {
		return utils.BadRequest(c, fmt.Errorf("captcha yang Anda berikan tidak valid"))
	}
	creds := models.ResetPasswordRequest{}
	err := c.BodyParser(&creds)
	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
		return utils.BadRequest(c, err)

	}
	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	at, err := db.Auth(creds.Email)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jika email yang Anda masukkan benar, maka Anda akan menerima email berisikan kata sandi sementara",
		})
	}

	newPassword := randstr.String(16)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "a")
	mailer.SetHeader("To", creds.Email)
	mailer.SetHeader("Subject", "Reset Password")
	mailer.SetBody("text/html", "Kata sandi Anda telah direset, kata sandi sementara Anda adalah: "+newPassword)
	d := gomail.NewDialer("a", 587, "a", "a")
	if err := d.DialAndSend(mailer); err != nil {
		return utils.InternalServerError(c, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 8)
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	if err != nil {
		return utils.InternalServerError(c, err)
	}
	body := &models.User{
		ID:        at.ID,
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
	}

	err = db.UpdateUserPassword(body)
	if err != nil {
		return utils.InternalServerError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Jika email yang Anda masukkan benar, maka Anda akan menerima email berisikan kata sandi sementara",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	body := models.RefreshToken{}
	err := c.BodyParser(&body)

	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return utils.BadRequest(c, err)
	}

	redisDb, err := database.RedisConnect()

	if err != nil {
		return utils.InternalServerError(c, err)
	}

	claims, err := utils.ExtractRefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Token tidak valid",
			"note":    err.Error(),
		})
	}

	deleted, err := redisDb.DeleteAuth(claims.RefreshUuid)

	if err != nil || deleted == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Token tidak valid",
			"note":    err.Error(),
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
		return utils.InternalServerError(c, err)
	}

	refresh := redisDb.CreateAuth(claims.User, token)
	if refresh != nil {
		return utils.InternalServerError(c, err)
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
