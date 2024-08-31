package user

import (
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) Create(c *fiber.Ctx) error {

	err := h.userService.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "damn thing failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "user service up and running.",
	})
}
