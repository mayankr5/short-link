package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/handler"
	"github.com/mayankr5/url_shortner/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handler.Login)

	app.Post("/create-short-url", middleware.Authentication, handler.CreateShortUrl)
	app.Get("/:shortUrl", handler.HandleShortUrlRedirect)
}
