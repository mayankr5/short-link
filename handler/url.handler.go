package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/config"
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
			"message": "incorrect time format",
			"error":   err,
		})
	}

	if cacheDuration.Before(time.Now()) {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "elapsed time",
			"error":   errors.New("please provide future time from current time"),
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
	host := config.Config("RAILWAY_PUBLIC_DOMAIN")
	if host == "" {
		host = "http://0.0.0.0:3000/"
	} else {
		host = "https://" + host + "/"
	}

	userURL := model.UserURL{
		ID:          uuid.New(),
		OriginalURL: creationRequest.OriginalURL,
		ShortURL:    host + shortUrl,
		UserID:      creationRequest.UserId,
		Validity:    cacheDuration,
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
		Created_at  time.Time `json:"created_at"`
		Validity    time.Time `json:"validity"`
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
			Created_at:  userURL.CreatedAt,
			Validity:    userURL.Validity,
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
	host := config.Config("RAILWAY_PUBLIC_DOMAIN")
	if host == "" {
		host = "http://0.0.0.0:3000/"
	} else {
		host = "https://" + host + "/"
	}

	var userUrl model.UserURL
	store.DB.Db.Where("short_url = ?", host+shortUrl).First(&userUrl)

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
