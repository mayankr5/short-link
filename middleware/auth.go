package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
	"gorm.io/gorm"
)

func getToken(accessToken string) (*model.AuthToken, error) {
	auth_token := new(model.AuthToken)
	if err := store.DB.Db.Where("token = ?", accessToken).First(&auth_token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return auth_token, nil
}

func Authentication(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")

	if accessToken == "" {
		accessToken = c.Get("Authorization", "")
		if accessToken != "" {
			accessToken = strings.TrimPrefix(accessToken, "Bearer ")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "missing token",
				"error":   "token not found",
			})
		}
	}

	auth_token, err := getToken(accessToken)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err,
		})
	} else if auth_token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid token",
			"data":    nil,
		})
	}

	c.Locals("auth_token", *auth_token)

	return c.Next()
}
