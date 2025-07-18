package service

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
)

type PostServiceInterface interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
}
