package controllers

import (
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/platform/database"

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
func Login(c *fiber.Ctx) error{
	creds := models.Credentials{}
	err := c.BodyParser(&creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}
	
	validate := utils.NewValidator()
	if err := validate.Struct(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": "cannot connect to database",
		}) 
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password),8)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": "cannot hash password",
		}) 
	}

	auth, err := db.Auth(creds.Username)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"message": "wrong username",
		}) 
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(auth.Password)); err == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"message": "wrong password",
		}) 
	}

	token, err := utils.GenerateNewAccessToken(auth.ID, auth.Email, auth.Level)
	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg": "cannot create jwt token",
        })
    }
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}