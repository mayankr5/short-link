package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func insertToken(auth_token *model.AuthToken) error {
	db := store.DB.Db

	if err := db.Create(&auth_token).Error; err != nil {
		return err
	}
	return nil
}

func GenerateToken(user User) (string, error) {

	s := user.ID.String() + user.Username + user.Email + time.Now().Format("2006-01-02 15:04:05")
	h := sha1.New()
	h.Write([]byte(s))
	tokenString := hex.EncodeToString(h.Sum(nil))

	auth_token := model.AuthToken{
		ID:     uuid.New(),
		Token:  tokenString,
		UserID: user.ID,
	}

	if err := insertToken(&auth_token); err != nil {
		return "", err
	}

	return auth_token.Token, nil
}
