package handler

import "github.com/gofiber/fiber/v2"

type APIHandler interface {
	RegisterRoutes(fiber.Router)
}
