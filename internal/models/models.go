package models

import (
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
