package handler

import (
	"blogio/internal/handler/responses"
	Uservice "blogio/internal/service/interfaces"
	"github.com/gofiber/fiber/v2"
)

// type UserHandler interface {
// 	FindAll(fiber.Ctx) ([]entity.User, error)
// 	FindByID(fiber.Ctx) (entity.User, error)
// }

type UserHandler struct {
	userService Uservice.UserInterface
}

func NewUserHandler(userService Uservice.UserInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) FindAll(c *fiber.Ctx) error{
	user, err := handler.userService.FindAll(c.Context())

	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Failed to fetch users",
			Data:  err.Error(),
		})
	}

	return  c.Status(fiber.StatusOK).JSON(responses.Response{
		Status: fiber.StatusOK,
		Message: "Berhasil fetch data",
		Data: user,
	})
}

func (handler *UserHandler) FindByID(c *fiber.Ctx) error{
	 id := c.Params("id")
	//  hex_id, _ := primitive.ObjectIDFromHex(id)
	 user, err := handler.userService.FindByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
				Status: fiber.StatusInternalServerError,
				Message: "Failed to fetch user",
				Data:  err.Error(),
			})
		}
		return  c.Status(fiber.StatusOK).JSON(responses.Response{
			Status: fiber.StatusOK,
			Message: "Berhasil fetch data",
			Data: user,
		})
}