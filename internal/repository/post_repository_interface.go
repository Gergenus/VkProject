package repository

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
)

type PostRepositoryInterface interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
}
