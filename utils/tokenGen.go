package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func GenerateToken(user User) string {

	// Need To add secret key in this
	s := user.ID.String() + user.Username + user.Email + time.Now().Format("2006-01-02 15:04:05")
	h := sha1.New()
	h.Write([]byte(s))
	tokenString := hex.EncodeToString(h.Sum(nil))

	return tokenString
}
