package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	PasswordHash string    `json:"password_hash"`
}

type ProductPost struct {
	ID           int       `json:"id,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty"`
	Subject      string    `json:"subject"`
	PostText     string    `json:"post_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
