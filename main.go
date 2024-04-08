package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/routes"
	"github.com/mayankr5/url_shortner/store"
)

func main() {
	fmt.Printf("hello")
	app := fiber.New()

	routes.SetupRoutes(app)

	store.InitializeStore()
	store.Connect()
	app.Listen(":3000")
}
