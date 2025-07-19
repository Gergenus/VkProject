package repository

import (
	"context"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/transport/http/dto"
)

type PostRepositoryInterface interface {
	CreatePost(ctx context.Context, post models.ProductPost) (int, error)
	Posts(ctx context.Context, page, pageSize int, userId, sortBy, sortDir string, minPrice, maxPrice float64) (*[]dto.ResponsePost, error)
}
