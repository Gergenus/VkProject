package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Login        string
	PasswordHash string
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Post struct {
	ID           int       `json:"id,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty"`
	Subject      string    `json:"subject"`
	PostText     string    `json:"post_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type ReturnPost struct {
	ID           int       `json:"id,omitempty"`
	Login        string    `json:"login"`
	Subject      string    `json:"subject"`
	PostText     string    `json:"post_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	IsOwner      bool      `json:"is_owner"`
}
