package utils

import "github.com/gofiber/fiber/v2"

func Conflict(c *fiber.Ctx, e error, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   true,
		"message": message + " telah digunakan",
		"note":    e,
	})
}

func BadRequest(c *fiber.Ctx, e error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   true,
		"message": "Mohon cek kembali data yang Anda masukkan",
		"note":    e.Error(),
	})
}

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

func NotFound(c *fiber.Ctx, e error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   true,
		"message": "Data tidak ditemukan",
		"note":    e.Error(),
	})
}
