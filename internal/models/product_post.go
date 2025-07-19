package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductPost struct {
	ID           int       `json:"id,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty"`
	Subject      string    `json:"subject"`
	PostText     string    `json:"post_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
