package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/store"
	"github.com/mayankr5/url_shortner/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type AuthToken struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token  string    `gorm:"unique" json:"token"`
	UserID uuid.UUID `json:"user_id"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func Login(c *fiber.Ctx) error {
	var userCredential LoginRequest
	if err := c.BodyParser(&userCredential); err != nil {
		return c.JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	var user User

	result := store.DB.Db.Where("Username = ? AND Password = ?", userCredential.Username, userCredential.Password).Find(&user)
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

	userRes := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	token := utils.GenerateToken(utils.User(userRes))

	auth_token := AuthToken{
		ID:     uuid.New(),
		Token:  token,
		UserID: user.ID,
	}

	if err := store.DB.Db.Create(&auth_token); err.Error != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusNotImplemented,
			"error":  err.Error,
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "user login",
		"token":   token,
		"user":    userRes,
	})
}

func Signup(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	user.ID = uuid.New()

	if err := store.DB.Db.Create(&user); err.Error != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error,
		})
	}
	userRes := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	token := utils.GenerateToken(utils.User(userRes))

	auth_token := AuthToken{
		ID:     uuid.New(),
		Token:  token,
		UserID: user.ID,
	}

	if err := store.DB.Db.Create(&auth_token); err.Error != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotImplemented,
			"message": "internal error",
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "user registered",
		"token":   token,
		"user":    userRes,
	})
}

func Logout(c *fiber.Ctx) error {
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

	var auth_token AuthToken
	result := store.DB.Db.Where("token = ?", accessToken).Delete(&auth_token)

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

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "user logout",
	})
}
