package handler

import (
	"blogio/helper"
	"blogio/internal/domain/entity"
	"blogio/internal/handler/responses"
	Uservice "blogio/internal/service/interfaces"
	"fmt"
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

//var validate = validator.New()

func (handler *UserHandler) CreateUser (c *fiber.Ctx) error {
	
	Validator := helper.NewValidator()

	// user := new(entity.User)

	var user entity.User

	if err := c.BodyParser(&user); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse{
			Error:       true,
			FailedField: "Body Parsing",
			Tag:         "parse_error",
			Value:       err.Error(),
		})
	}

	if errs := Validator.Validator(user); len(errs)> 0 {
		errorMessage := make([]string,0)

		for _, err := range errs {
			errorMessage = append(errorMessage,fmt.Sprintf(
				"Error on field %s, condition %s, value %s",
				err.FailedField,
				err.Tag,
				err.Value,
			) )
		}

	
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse{
			Error:       true,
			FailedField: "Validation",
			Tag:         "validation_error",
			Value:       errorMessage,
		})
	}
	
	users, err := handler.userService.CreateUser(c.Context(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Failed to create user",
			Data:  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status: fiber.StatusOK,
		Message: "Berhasil Create data",
		Data: users,
	})
	
}


