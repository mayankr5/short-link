package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mayankr5/url_shortner/handler"
	"github.com/mayankr5/url_shortner/middleware"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Authorisation APIs
	user_api := api.Group("/auth")
	user_api.Post("/login", handler.Login)
	user_api.Post("/signup", handler.Signup)

	//Protected APIs
	url_api := api.Group("/url", middleware.Authentication)
	url_api.Post("/create-short-url", handler.CreateShortUrl)
	url_api.Get("/get-urls", handler.GetURLs)
	url_api.Get("/logout", handler.Logout)

	// Public APIs
	app.Get("/:shortUrl", handler.HandleShortUrlRedirect)
}
