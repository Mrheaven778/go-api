package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrheaven778/go-astro-crud/controllers"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")
	api.Post("/login", controllers.LoginHandler)
	api.Post("/register", controllers.RegisterHandler)
}
