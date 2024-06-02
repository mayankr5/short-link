package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
)

type UserResp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	TotalUrl int64  `json:"total_urls"`
}

func GetUser(c *fiber.Ctx) error {

	userId := c.Params("user_id")

	var user model.User

	if err := store.DB.Db.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(&fiber.Map{
			"status": "user not found",
			"error":  err,
			"data":   nil,
		})
	}

	var userURLs []model.UserURL

	store.DB.Db.Where("user_id = ?", userId).Find(&userURLs)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "user details",
		"error":  nil,
		"data": &fiber.Map{
			"user": &UserResp{
				Name:     user.Name,
				Email:    user.Email,
				Username: user.Username,
				TotalUrl: int64(len(userURLs)),
			},
		},
	})
}

func UpdateUser(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "update user",
		"error":  nil,
		"data": &fiber.Map{
			"user": &UserResp{},
		},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "delete user",
		"error":  nil,
		"data":   nil,
	})
}
