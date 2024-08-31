package healthcheck

import "github.com/gofiber/fiber/v2"

type HealthHandler struct {
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Service up and running.",
	})
}

func (h *HealthHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.HealthCheck)
}

func NewHealthCheckHandler() *HealthHandler {
	return &HealthHandler{}
}
