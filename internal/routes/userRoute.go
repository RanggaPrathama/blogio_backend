package routes

import (
	"blogio/internal/handler"

	"github.com/gofiber/fiber/v2"
)



func UserRoute(app *fiber.App, handler *handler.UserHandler)  {
	app.Get("/users", handler.FindAll )

}