package routes

import (
	"blogio/internal/handler"

	"github.com/gofiber/fiber/v2"
)



func UserRoute(app *fiber.App, handler *handler.UserHandler)  {
	api := app.Group("/api")
	user := api.Group("/users")
	user.Get("/", handler.FindAll )
	user.Get("/:id", handler.FindByID)
	user.Post("/create", handler.CreateUser)
	user.Put("/update/:id", handler.UpdateUser)
	user.Delete("/delete/:id", handler.DeleteUser)
}

