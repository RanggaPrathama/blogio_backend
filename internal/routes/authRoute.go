package routes

import (
	"blogio/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, handle *handler.AuthHandler){
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/login", handle.Login)
}