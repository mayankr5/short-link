package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func Login(c *fiber.Ctx) error {
	var userCredential LoginRequest
	if err := c.BodyParser(&userCredential); err != nil {
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": 400,
		})
	}

	// verify user and get user detail
	var user User

	token := utils.GenerateToken(utils.User(user))
	// Save token with user detail

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "user login",
		"token":   token,
	})
}
