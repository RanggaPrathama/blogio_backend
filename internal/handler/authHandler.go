package handler

import (
	"blogio/internal/service/interfaces"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService interfaces.AuthInterface
}

func NewAuthHandler(authService interfaces.AuthInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}


func (handle *AuthHandler) Login(c *fiber.Ctx) error {

	var request struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Failed to parse body",
			"data":    err.Error(),
		})
	}

	user, err := handle.AuthService.Login(c.Context(), request.Email, request.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to login",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil login",
		"data":    user,
	})
}