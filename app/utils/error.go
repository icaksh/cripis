package utils

import "github.com/gofiber/fiber/v2"

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "Anda tidak diperkenankan melakukan aksi ini",
	})
}

func InternalServerError(c *fiber.Ctx, e error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   true,
		"message": "Terjadi kesalahan (Internal Server Error)",
		"note":    e.Error(),
	})
}
