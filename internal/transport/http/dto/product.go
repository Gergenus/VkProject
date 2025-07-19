package dto

import "time"

type ResponsePost struct {
	ID           int       `json:"id,omitempty"`
	Login        string    `json:"login"`
	Subject      string    `json:"subject"`
	PostText     string    `json:"post_text"`
	ImageAddress string    `json:"image_address"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	IsOwner      bool      `json:"is_owner"`
}
