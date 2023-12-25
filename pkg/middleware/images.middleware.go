package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icaksh/cripis/app/utils"
)

func AccessControl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		folderUUID := c.Params("uuid")
		at, err := utils.ExtractTokenMetadata(c)
		if err != nil || at.Role != 1 || at.User.String() != folderUUID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   true,
				"message": "Anda tidak diperkenankan melakukan aksi ini",
			})
		}
		return c.Next()
	}
}
