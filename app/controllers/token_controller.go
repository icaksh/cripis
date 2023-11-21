package controllers

import (
	"github.com/icaksh/cripis/app/utils"

	"github.com/gofiber/fiber/v2"
)

// RefreshToken method for refresh token.
// @Description Refresh token.
// @Summary refresh token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/refresh [get]
func RefreshToken(c *fiber.Ctx) error {
    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }
	token, err := utils.GenerateNewAccessToken(claims.User, claims.Email, claims.Role)

	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	}) 
}