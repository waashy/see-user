package healthcheck

import "github.com/gofiber/fiber/v2"

type HealthCheckhandler struct {
}

func (h *HealthCheckhandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Service up and running.",
	})
}

func (h *HealthCheckhandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.HealthCheck)
}

func NewHealthCheckHandler() *HealthCheckhandler {
	return &HealthCheckhandler{}
}
