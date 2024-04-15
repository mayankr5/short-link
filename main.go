package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mayankr5/url_shortner/routes"
	"github.com/mayankr5/url_shortner/store"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./public")

	routes.SetupRoutes(app)

	store.InitializeStore()
	store.Connect()

	app.Listen(":3000")
}
