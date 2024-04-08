package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/store"
	"github.com/mayankr5/url_shortner/utils"
)

type UrlCreationRequest struct {
	OriginalURL string    `json:"original_url" binding:"required"`
	UserId      uuid.UUID `json:"user_id" binding:"required"`
}

type UserURLs struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OriginalURL string    `gorm:"unique" json:"original_url"`
	ShortURL    string    `gorm:"unique" json:"short_url"`
	Visiter     int       `json:"visiter"`
	UserID      uuid.UUID `json:"user_id"`
}

func GetUrls(c *fiber.Ctx) error {
	return nil
}

func CreateShortUrl(c *fiber.Ctx) error {
	var creationRequest UrlCreationRequest
	if err := c.BodyParser(&creationRequest); err != nil {
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}

	shortUrl := utils.GenerateShortLink(creationRequest.OriginalURL, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.OriginalURL, creationRequest.UserId)

	userURL := UserURLs{
		ID:          uuid.New(),
		OriginalURL: creationRequest.OriginalURL,
		ShortURL:    shortUrl,
		UserID:      creationRequest.UserId,
	}

	result := store.DB.Db.Create(&userURL)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": result.Error,
		})
	}

	host := "http://localhost:3000/"
	return c.JSON(fiber.Map{
		"status":    fiber.StatusOK,
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	}, "application/problem+json")

}

func HandleShortUrlRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	var userUrl UserURLs
	result := store.DB.Db.Where("shorturl = ?", shortUrl).Find(&userUrl)

	userUrl.Visiter = userUrl.Visiter + 1

	store.DB.Db.Save(&userUrl)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": result.Error,
		})
	}
	return c.Redirect(initialUrl)
}
