package handler

import (
	"errors"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
	"github.com/mayankr5/url_shortner/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identity string `json:"Identity"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*model.User, error) {
	db := store.DB.Db
	var user model.User
	if err := db.Where(&model.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(u string) (*model.User, error) {
	db := store.DB.Db
	var user model.User
	if err := db.Where(&model.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx) error {
	var userCredential LoginRequest
	if err := c.BodyParser(&userCredential); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "error on login request",
			"error":   err.Error(),
		})
	}

	var userRes UserResponse
	userModel, err := new(model.User), *new(error)

	if valid(userCredential.Identity) {
		userModel, err = getUserByEmail(userCredential.Identity)
	} else {
		userModel, err = getUserByUsername(userCredential.Identity)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err,
		})
	} else if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid identity or password",
			"data":    nil,
		})
	} else {
		userRes = UserResponse{
			ID:       userModel.ID,
			Name:     userModel.Name,
			Username: userModel.Username,
			Email:    userModel.Email,
		}
	}

	if !CheckPasswordHash(userCredential.Password, userModel.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid identity or password",
			"data":    nil,
		})
	}

	token, err := utils.GenerateToken(utils.User(userRes))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "interanl server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user login",
		"data":    fiber.Map{"user": userRes, "token": token},
	})
}

// add check whether a username or email is present in database already
func Signup(c *fiber.Ctx) error {

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "error on signup request",
			"error":   err.Error(),
		})
	}

	userModel, err := new(model.User), *new(error)
	userModel, err = getUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err,
		})
	} else if userModel != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "duplicate user error",
			"error":   "email is already present",
		})
	}
	userModel, err = getUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err,
		})
	} else if userModel != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "duplicate user error",
			"error":   "username is already present",
		})
	}

	hash_pass, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	user.ID = uuid.New()
	user.Password = hash_pass

	if err := store.DB.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "error on creating user",
			"error":   err.Error(),
		})
	}

	userRes := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	token, err := utils.GenerateToken(utils.User(userRes))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "error on token generation",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user registered",
		"data": fiber.Map{
			"user":  userRes,
			"token": token,
		},
	})
}

func Logout(c *fiber.Ctx) error {

	auth_token := c.Locals("auth_token").(model.AuthToken)
	store.DB.Db.Delete(&auth_token)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user logout",
		"data":    nil,
	})
}
