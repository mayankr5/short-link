package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserURLs struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Visiter     int       `json:"visiter"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `gorm:"foreignKey:UserID"`
}

type AuthToken struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token     string    `gorm:"unique" json:"token"`
	UserID    uuid.UUID `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
