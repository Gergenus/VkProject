package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Gergenus/VkProject/internal/models"
)

func (u *UserService) CreatePost(ctx context.Context, post models.Post) (int, error) {
	const op = "post.service.CreatePost"
	log := u.log.With(slog.String("op", op))
	log.Info("creating post", slog.String("userID", post.UserID.String()))
	id, err := u.postRepo.CreatePost(ctx, post)
	if err != nil {
		log.Error("failed to create post", slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("post created", slog.Int("postId", id))
	return id, nil
}
