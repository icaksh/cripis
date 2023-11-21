package controllers

import (
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
func Register(c *fiber.Ctx) error{
	creds := models.SignUp{}
	users := models.User{}
	profile := models.UserProfile{}
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
			"message": err.Error(),
			"note": "cant connect to database",
		}) 
	}

	
	userId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password),8)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
			"note": "cant generating salt password",
		}) 
	}
	
	isUsernameExist := db.CheckDuplicate("username", creds.Username)
	
	if isUsernameExist {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"note": "username has been exist",
		})
	}

	isEmailUsed := db.CheckDuplicate("email", creds.Email)
	
	if isEmailUsed {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"note": "email has been used",
		})
	}

	profile = models.UserProfile{
		ID: userId,
		FirstName: creds.FirstName,
		LastName: creds.LastName,
	}

	users = models.User{
		ID : userId,
		Email: creds.Email,
		Username : creds.Username,
		Password : string(hashedPassword),
		CreatedAt : time.Now(),
		Level : 0,
		Status : 1,
	}

	if err := db.CreateUser(&users); err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
			"note": "cannot store user to database",
		})
	}

	if err := db.CreateUserProfile(&profile); err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
			"note": "cannot store profile to database",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(users)
}
