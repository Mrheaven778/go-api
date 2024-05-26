package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrheaven778/go-astro-crud/controllers"
)

func TaskRoutes(app *fiber.App) {
	api := app.Group("/api/tasks")
	api.Get("/", controllers.GetTasksHandler)
	api.Get("/:id", controllers.GetTaskHandler)
	api.Post("/:userId", controllers.CreateTaskHandler)
	api.Delete("/:userId/:id", controllers.DeleteTaskHandler)
}
