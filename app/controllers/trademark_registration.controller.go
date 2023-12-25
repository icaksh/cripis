package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"os"
	"strings"
	"time"
)

func GetTrademarkRegistrations(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Anda tidak diperkenankan melakukan aksi ini",
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

	query, err := db.GetTrademarkRegistrations(at.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Senarai pendaftaran merek dagang tidak ditemukan",
			"note":    "cannot get trademark registration, err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

func GetAllTrademarkRegistrations(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil || at.Role != 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Anda tidak diperkenankan melakukan aksi ini",
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Tidak dapat terkoneksi dengan database",
			"note":    "cant connect to database, err: " + err.Error(),
		})
	}

	query, err := db.AdGetTrademarkRegistrations()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Senarai pendaftaran merek dagang tidak ditemukan",
			"note":    "cannot get trademark registration, err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

func CreateTrademarkRegistration(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Anda tidak diperkenankan melakukan aksi ini",
		})
	}
	body := models.TrademarkRegistrationRequest{}
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat diparse, mohon cek kembali",
			"note":    "cannot parse data, err: " + err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat divalidasi, mohon cek kembali",
			"note":    "cannot validate data, err: " + err.Error(),
		})
	}

	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat divalidasi, mohon cek kembali",
			"note":    "cannot validate data, err: " + err.Error(),
		})
	}

	smeCertificate, err := c.FormFile("sme_certificate")
	isSmeExist := true
	if err != nil {
		isSmeExist = false
	}

	registerSignature, err := c.FormFile("register_signature")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Data tidak dapat divalidasi, mohon cek kembali",
			"note":    "cannot validate data, err: " + err.Error(),
		})
	}
	currentDate := time.Now()

	dateString := currentDate.Format("2006/01")
	dir := "./public/uploads/" + dateString + "/" + at.User.String()
	imageName := dir + "/" + image.Filename

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	registerSignatureName := dir + "/" + registerSignature.Filename

	err = c.SaveFile(image, imageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot save image, err: " + err.Error(),
		})
	}

	var smeCertificateName string
	if isSmeExist {
		smeCertificateName = dir + "/" + smeCertificate.Filename
		err = c.SaveFile(smeCertificate, smeCertificateName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Terjadi kesalahan (Internal Server Error)",
				"note":    "cannot save image, err: " + err.Error(),
			})
		}
	}

	err = c.SaveFile(registerSignature, registerSignatureName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot save image, err: " + err.Error(),
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot connect to database, err: " + err.Error(),
		})
	}

	trademarkId := uuid.New()

	registration := &models.TrademarkRegistration{
		ID:                trademarkId,
		CreatedAt:         time.Now(),
		CreatedBy:         at.User,
		SMECertificate:    smeCertificateName,
		RegisterSignature: registerSignatureName,
	}
	trademark := &models.Trademark{
		RegisterNumber: utils.GenerateRegistrationNumber(),
		TrademarkName:  strings.ToUpper(body.TrademarkName),
		Class:          body.Class,
		OwnerName:      strings.ToUpper(body.OwnerName),
		Address:        body.Address,
		Village:        body.Village,
		District:       body.District,
		Regency:        body.Regency,
		Province:       body.Province,
		Image:          imageName,
	}

	err = db.CreateTrademarkRegistration(registration, trademark)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Terjadi kesalahan (Internal Server Error)",
			"note":    "cannot create trademark registration, err: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan pendaftaran merek dagang",
	})
}

func GetTrademarkRegistration(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Anda tidak diperkenankan melakukan aksi ini",
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Tidak dapat terkoneksi dengan database",
			"note":    "cannot connect to database, err: " + err.Error(),
		})
	}

	id := c.Params("id")
	query, err := db.GetTrademarkRegistration(uuid.Must(uuid.Parse(id)), at.User)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Pendaftaran merek dagang tidak ditemukan",
			"note":    "cannot get trademark registration, err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(query)
}
