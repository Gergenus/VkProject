package service

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
)

type PostServiceInterface interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
	Posts(ctx context.Context, page, pageSize int, userId, sortBy, sortDir string, minPrice, maxPrice float64) (*[]models.ReturnPost, error)
}
