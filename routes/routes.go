package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/handler"
	"github.com/mayankr5/url_shortner/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Autherisation APIs
	app.Post("/login", handler.Login)
	app.Post("/signup", handler.Signup)
	app.Get("/logout", handler.Logout)

	//Protected APIs
	app.Post("/create-short-url", middleware.Authentication, handler.CreateShortUrl)
	app.Get("/get-urls", middleware.Authentication, handler.GetURLs)

	// Public APIs
	app.Get("/:shortUrl", handler.HandleShortUrlRedirect)
}
