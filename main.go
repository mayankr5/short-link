package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/handler"
	"github.com/mayankr5/url_shortner/store"
)

func main() {
	fmt.Printf("hello")
	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/create-short-url", handler.CreateShortUrl)

	app.Get("/:shortUrl", handler.HandleShortUrlRedirect)

	store.InitializeStore()
	app.Listen(":3000")
}
