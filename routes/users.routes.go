package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrheaven778/go-astro-crud/controllers"
	"github.com/mrheaven778/go-astro-crud/middleware"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/users")
	api.Get("/", middleware.JWTProtected(), controllers.GetUsers)
	api.Get("/:id", middleware.JWTProtected(), controllers.GetUser)
	api.Get("/:id", controllers.GetUser)
	api.Put("/:id", controllers.UpdateUser)
	api.Delete("/:id", controllers.DeleteUser)
}
