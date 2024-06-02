package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mayankr5/url_shortner/routes"
	"github.com/mayankr5/url_shortner/store"
)

func main() {
	app := fiber.New(fiber.Config{
		Network:      "tcp",
		ServerHeader: "Fiber",
		AppName:      "Short-Link v1.0.1",
	})

	app.Use(cors.New())

	app.Static("/", "./public", fiber.Static{
		CacheDuration: -1,
		MaxAge:        0, // Set to 0 to prevent caching
	})

	routes.SetupRoutes(app)

	store.InitializeStore()
	store.Connect()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
