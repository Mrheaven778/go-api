package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mrheaven778/go-astro-crud/config"
	"github.com/mrheaven778/go-astro-crud/db"
	"github.com/mrheaven778/go-astro-crud/routes"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()

	port := config.GetPort()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// Routes
	routes.UserRoutes(app)
	routes.TaskRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(":" + port)
	fmt.Print("Server is running on port: " + port)
}
