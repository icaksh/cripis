package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"
	"strconv"
	"strings"
	"time"
)

func CreateTrademark(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return utils.Unauthorized(c)
	}
	body := models.TrademarkRegistrationRequest{}
	err = c.BodyParser(&body)

	if err != nil {
		return utils.BadRequest(c, err)
	}

	validate := utils.NewValidator()
	if err := validate.Struct(&body); err != nil {
		return utils.BadRequest(c, err)
	}

	image, err := c.FormFile("image")
	if err != nil {
		return utils.BadRequest(c, err)
	}

	smeCertificate, err := c.FormFile("sme_certificate")
	isSmeExist := true
	if err != nil {
		isSmeExist = false
	}

	registerSignature, err := c.FormFile("register_signature")
	if err != nil {
		return utils.BadRequest(c, err)
	}

	imgUpload, imgMessage, err := utils.UploadFiletoS3(at.User.String(), image)
	if !imgUpload {
		fmt.Println(imgMessage)
		return utils.InternalServerError(c, err)
	}

	registerSignatureUpload, registerSignatureMessage, err := utils.UploadFiletoS3(at.User.String(), registerSignature)
	if !registerSignatureUpload {
		fmt.Println(registerSignatureMessage)
		return utils.InternalServerError(c, err)
	}

	var smeCertificateMessage string
	if isSmeExist {
		smeCertificateUpload, sme, err := utils.UploadFiletoS3(at.User.String(), smeCertificate)
		if !smeCertificateUpload {
			fmt.Println(smeCertificateMessage)
			return utils.InternalServerError(c, err)
		}
		smeCertificateMessage = sme
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	trademarkId := uuid.New()

	trademark := &models.Trademark{
		ID:                trademarkId,
		CreatedAt:         time.Now(),
		CreatedBy:         at.User,
		SMECertificate:    smeCertificateMessage,
		RegisterSignature: registerSignatureMessage,
		RegisterNumber:    utils.GenerateRegistrationNumber(),
		TrademarkName:     strings.ToUpper(body.TrademarkName),
		Class:             body.Class,
		OwnerName:         strings.ToUpper(body.OwnerName),
		Address:           body.Address,
		Village:           body.Village,
		District:          body.District,
		Regency:           body.Regency,
		Province:          body.Province,
		Image:             imgMessage,
	}

	err = db.CreateTrademark(trademark)
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

func GetTrademarks(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil || at.Role != 1 {
		return utils.Unauthorized(c)
	}
	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	res, err := db.GetTrademarks()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
			"note":    "cannot get trademark err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func GetTrademarksBySearch(c *fiber.Ctx) error {
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
	query, err = db.GetTrademarksByName(queryValue)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
			"note":    "cannot get trademark err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

func GetTrademarksByUser(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return utils.Unauthorized(c)
	}
	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
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

func GetTrademarkById(c *fiber.Ctx) error {
	_, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"note":    "Anda tidak diperkenankan melakukan aksi ini",
			"message": err.Error(),
		})
	}

	db, err := database.Connect()
	trademarkId := c.Params("id")

	res, err := db.GetTrademark(trademarkId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
			"note":    "cannot get trademark err:" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func UpdateTrademark(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return utils.Unauthorized(c)
	}

	body := models.TrademarkEditRequest{}
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

	trademarkId := body.ID.String()
	fmt.Println(trademarkId)
	trademark, err := db.GetTrademark(trademarkId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
			"note":    "cannot get trademark err:" + err.Error(),
		})
	}

	if at.User != trademark.CreatedBy || at.Role != 1 {
		return utils.Unauthorized(c)
	}

	image, err := c.FormFile("new_image")
	isImageUpdated := true
	if err != nil {
		isImageUpdated = false
	}

	smeCertificate, err := c.FormFile("new_sme_certificate")
	isSmeExist := true
	if err != nil {
		isSmeExist = false
	}

	registerSignature, err := c.FormFile("new_register_signature")
	isRegisterSignatureUpdated := true
	if err != nil {
		isRegisterSignatureUpdated = false
	}

	fmt.Println(uuid.MustParse(trademarkId))
	data := &models.Trademark{
		ID:                uuid.MustParse(trademarkId),
		TrademarkName:     strings.ToUpper(body.TrademarkName),
		Class:             body.Class,
		OwnerName:         strings.ToUpper(body.OwnerName),
		Address:           body.Address,
		Village:           body.Village,
		District:          body.District,
		Regency:           body.Regency,
		Province:          body.Province,
		Image:             body.Image,
		SMECertificate:    body.SMECertificate,
		RegisterSignature: body.RegisterSignature,
		UpdatedAt:         time.Now(),
		Status:            1,
	}
	if isImageUpdated {
		imgUpload, imgMessage, err := utils.UploadFiletoS3(at.User.String(), image)
		if !imgUpload {
			fmt.Println(imgMessage)
			return utils.InternalServerError(c, err)
		}
		data.Image = imgMessage
	}

	if isRegisterSignatureUpdated {
		registerSignatureUpload, registerSignatureMessage, err := utils.UploadFiletoS3(at.User.String(), registerSignature)
		if !registerSignatureUpload {
			fmt.Println(registerSignatureMessage)
			return utils.InternalServerError(c, err)
		}
		data.RegisterSignature = registerSignatureMessage
	}

	if isSmeExist {
		smeCertificateUpload, smeCertificateMessage, err := utils.UploadFiletoS3(at.User.String(), smeCertificate)
		if !smeCertificateUpload {
			fmt.Println(smeCertificateMessage)
			return utils.InternalServerError(c, err)
		}
		data.SMECertificate = smeCertificateMessage
	}

	db, err = database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}
	err = db.UpdateTrademark(data)
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	db.CreateLog(at.User, "update trademark, id: "+trademarkId+" name: "+trademark.TrademarkName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengubah data",
	})
}

func UpdateTrademarkStatus(c *fiber.Ctx) error {
	au, err := utils.ExtractTokenMetadata(c)
	if err != nil || au.Role != 1 {
		return utils.Unauthorized(c)
	}

	body := models.TrademarkStatusRequest{}
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

	data := &models.Trademark{
		ID: uuid.MustParse(body.ID),
		Notes: sql.NullString{
			String: body.Notes,
			Valid:  true,
		},
		Status: body.Status,
	}

	expired := time.Now().Add(10 * 365 * 24 * time.Hour)
	if body.Status == 2 {
		data = &models.Trademark{
			ID: uuid.MustParse(body.ID),
			ApprovedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			ExpiredAt: sql.NullTime{
				Time:  expired,
				Valid: true},
			ApprovedBy: au.User,
			Notes: sql.NullString{
				String: body.Notes,
				Valid:  true,
			},
			Status: body.Status,
		}

		trademark, err := db.GetTrademark(body.ID)
		if err != nil {
			return utils.InternalServerError(c, err)
		}
		pdf, err := utils.CreateCertificate(&trademark)
		if err != nil {
			return utils.InternalServerError(c, err)
		}

		certificate, certificateMessage, err := utils.UploadFiletoS3PDF(au.User.String(), pdf, &trademark)
		if !certificate {
			return utils.InternalServerError(c, err)
		}
		fmt.Println(certificateMessage)
		data.File = sql.NullString{
			String: certificateMessage,
			Valid:  true,
		}
	}

	err = db.UpdateTrademarkStatus(data)

	if err != nil {
		return utils.InternalServerError(c, err)
	}

	db.CreateLog(au.User, "update trademark status, id: "+body.ID+" status: "+strconv.Itoa(body.Status))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Status merek dagang berhasil diubah",
	})
}

func DeleteTrademark(c *fiber.Ctx) error {
	at, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return utils.Unauthorized(c)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	id := c.Params("id")
	trademark, err := db.GetTrademark(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Merek dagang tidak ditemukan",
		})
	}

	if at.User != trademark.CreatedBy && at.Role != 1 {
		return utils.Unauthorized(c)
	}

	err = db.DeleteTrademark(id)
	if err != nil {
		return utils.InternalServerError(c, err)
	}

	db.CreateLog(at.User, "delete trademark, id: "+id+" name: "+trademark.TrademarkName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Merek dagang berhasil dihapus",
	})
}
