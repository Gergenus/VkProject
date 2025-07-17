package service

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
)

type UserServiceInterface interface {
	RegisterNewUser(ctx context.Context, login, password string) (models.User, error)
	Login(ctx context.Context, login, password string) (string, error)
}
