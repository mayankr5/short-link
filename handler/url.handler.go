package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
	"github.com/mayankr5/url_shortner/utils"
)

type UrlCreationRequest struct {
	OriginalURL    string    `json:"original_url" binding:"required"`
	UserId         uuid.UUID `json:"user_id"`
	ExpirationDate string    `json:"expiration_date"`
}

func CreateShortUrl(c *fiber.Ctx) error {
	var creationRequest UrlCreationRequest
	if err := c.BodyParser(&creationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "error on createShortUrl",
			"error":   err.Error(),
		})
	}
	creationRequest.UserId = c.Locals("auth_token").(model.AuthToken).UserID
	cacheDuration, err := time.Parse(time.RFC3339, creationRequest.ExpirationDate)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Send Time in correct format",
			"error":   err,
		})
	}

	shortUrl := utils.GenerateShortLink(creationRequest.OriginalURL, creationRequest.UserId)
	if err := store.SaveUrlMapping(shortUrl, creationRequest.OriginalURL, creationRequest.UserId, cacheDuration); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "error on redis storage",
			"error":   err.Error(),
		})
	}

	host := "http://localhost:3000/"

	userURL := model.UserURL{
		ID:          uuid.New(),
		OriginalURL: creationRequest.OriginalURL,
		ShortURL:    host + shortUrl,
		UserID:      creationRequest.UserId,
	}

	if err := store.DB.Db.Create(&userURL).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "shortUrl created successfully",
		"data":    fiber.Map{"short_url": host + shortUrl},
	})
}

func GetURLs(c *fiber.Ctx) error {

	type ResponseData struct {
		ID          uuid.UUID `json:"id"`
		OriginalURL string    `json:"original_url"`
		ShortURL    string    `json:"short_url"`
		Visiter     int       `json:"visiter"`
	}
	var userURLs []model.UserURL
	var responseData []ResponseData

	auth_token := c.Locals("auth_token").(model.AuthToken)

	store.DB.Db.Where("user_id = ?", auth_token.UserID).Find(&userURLs)

	for _, userURL := range userURLs {
		res := ResponseData{
			ID:          userURL.ID,
			OriginalURL: userURL.OriginalURL,
			ShortURL:    userURL.ShortURL,
			Visiter:     userURL.Visiter,
		}
		responseData = append(responseData, res)
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user URL's",
		"data":    responseData,
	})
}

func HandleShortUrlRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "URL is invalid or closed",
			"error":   err.Error(),
		})
	}

	var userUrl model.UserURL
	store.DB.Db.Where("short_url = ?", "http://localhost:3000/"+shortUrl).First(&userUrl)

	if userUrl.ShortURL == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "url not found",
			"error":   "url not found",
		})
	}

	userUrl.Visiter = userUrl.Visiter + 1
	store.DB.Db.Save(&userUrl)

	return c.Redirect(initialUrl)
}
