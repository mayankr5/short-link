package utils

import (
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/model"
	"github.com/mayankr5/url_shortner/store"
)

func VerifyToken(token string) (uuid.UUID, error) {

	var tokenFound model.AuthToken
	result := store.DB.Db.Where("token = ?", token).First(&tokenFound)

	if result.RowsAffected != 0 {
		return tokenFound.UserID, nil
	}

	return uuid.Nil, result.Error
}
