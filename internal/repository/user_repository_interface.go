package repository

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
)

type UserRepositoryInterface interface {
	SaveUser(ctx context.Context, login string, passwordHash string) (models.User, error)
	User(ctx context.Context, login string) (models.User, error)
}
