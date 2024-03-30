package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/store"
	"github.com/mayankr5/url_shortner/utils"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *fiber.Ctx) error {
	var creationRequest UrlCreationRequest
	if err := c.BodyParser(&creationRequest); err != nil {
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": 400,
		})
	}

	shortUrl := utils.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:3000/"
	return c.JSON(fiber.Map{
		"status":    200,
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	}, "application/problem+json")

}

func HandleShortUrlRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	return c.Redirect(initialUrl)
}
