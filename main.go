package main

import (
	"blogio/config"
	"blogio/internal/domain/repository"
	"blogio/internal/handler"
	"blogio/internal/routes"
	"blogio/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    config.MongoConnect()

    app.Get("/", func(c *fiber.Ctx) error {
       return c.SendString("Hello, World!")
    })

    // User
    repo := repository.NewUserRepository() 
    service := service.NewUserService(repo)
    handler := handler.NewUserHandler(service)   
    routes.UserRoute(app, handler)

    
    
    app.Listen(config.LoadEnv("PORT"))
}