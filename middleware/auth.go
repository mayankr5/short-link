package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/utils"
)

func Authentication(c *fiber.Ctx) error {
	accessToken := c.Cookies("accessToken")

	if accessToken == "" {
		accessToken = c.Get("Authorization", "")
		if accessToken != "" {
			accessToken = strings.TrimPrefix(accessToken, "Bearer ")
		} else {
			return c.JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Missing token",
			})
		}
	}

	_, err := utils.VerifyToken(accessToken)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid token",
		})
	}

	return c.Next()
}
