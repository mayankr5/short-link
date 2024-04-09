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
	ID          uuid.UUID `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Visiter     int       `json:"visiter"`
	UserID      uuid.UUID `json:"user_id"`
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

	host := "http://localhost:3000/"

	userURL := UserURLs{
		ID:          uuid.New(),
		OriginalURL: creationRequest.OriginalURL,
		ShortURL:    host + shortUrl,
		UserID:      creationRequest.UserId,
	}

	result := store.DB.Db.Create(&userURL)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": result.Error,
		})
	}

	return c.JSON(fiber.Map{
		"status":    fiber.StatusOK,
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	}, "application/problem+json")

}

func GetURLs(c *fiber.Ctx) error {
	var userURLs []UserURLs
	result := store.DB.Db.Where("user_id = ?", c.Locals("userID")).Find(&userURLs)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusNotImplemented,
			"error":  result.Error,
		})
	} else if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"error":  result.Error,
		})
	}

	return c.JSON(userURLs)
}

func HandleShortUrlRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Link is Invalid",
			"error":   err,
		})
	}

	var userUrl UserURLs
	result := store.DB.Db.Where("short_url = ?", "http://localhost:3000/"+shortUrl).First(&userUrl)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": result.Error,
		})
	}

	userUrl.Visiter = userUrl.Visiter + 1
	result = store.DB.Db.Save(&userUrl)

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": result.Error,
		})
	}
	return c.Redirect(initialUrl)
}
