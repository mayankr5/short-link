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
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/signup", handler.Signup)

	// Users APIs
	users := api.Group("/users/:user_id", middleware.Authentication)
	users.Get("/", handler.GetUser)
	users.Put("/update", handler.UpdateUser)
	users.Delete("/delete-account", handler.DeleteUser)

	//Protected APIs
	urls := api.Group("/urls", middleware.Authentication)
	urls.Post("/create-short-url", handler.CreateShortUrl)
	urls.Get("/get-urls", handler.GetURLs)
	urls.Get("/logout", handler.Logout)

	// Public APIs
	app.Get("/:shortUrl", handler.HandleShortUrlRedirect)
}
