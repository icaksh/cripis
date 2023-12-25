package controllers

//func ImageUpload(c *fiber.Ctx) error {
//	at, err := utils.ExtractTokenMetadata(c)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//			"error":   true,
//			"message": "Anda tidak diperkenankan melakukan aksi ini",
//		})
//	}
//
//	file, err := c.FormFile("image")
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error":   true,
//			"message": "Gagal mengupload gambar",
//			"note":    err.Error(),
//		})
//	}
//
//}
