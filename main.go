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
    serviceUser := service.NewUserService(repo)
    handlerUser := handler.NewUserHandler(serviceUser)   
    routes.UserRoute(app, handlerUser)

    // Auth 
    serviceAuth := service.NewAuthService(repo)
    handlerAuth := handler.NewAuthHandler(serviceAuth)
    routes.AuthRoute(app, handlerAuth)
    
    
    app.Listen(config.LoadEnv("PORT"))
}