package user

import (
	"github.com/gofiber/fiber/v2"
	userservice "github.com/waashy/see-user/service/user"
)

type UserHandler struct {
	userService userservice.UserService
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.Create)
}

func NewUserHandler(userService userservice.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
