package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	PasswordHash string    `json:"password_hash"`
}
