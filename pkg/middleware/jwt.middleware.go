package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/utils"
)

func JWTProtected(c *fiber.Ctx) error {
	err := utils.TokenValid(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	c.Next()
	return nil
}
