package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/transport/http/dto"
)

var (
	ErrIncorrectContents     = errors.New("contents should be less than 2500 chars")
	ErrIncorrectSubject      = errors.New("subject should be less than 100 chars")
	ErrIncorrectPrice        = errors.New("subject cannot be negative")
	ErrIncorrectImageAddress = errors.New("invalid photo type")
	ErrHeadRequestFailed     = errors.New("failed to get file size")
	ErrIncorrectImageSize    = errors.New("content size should be less than 5mb")
)

var allowedPhotoTypes = map[string]bool{
	"jpg": true,
	"png": true,
}

func (u *UserService) CreatePost(ctx context.Context, post models.ProductPost) (int, error) {
	const op = "post.service.CreatePost"
	log := u.log.With(slog.String("op", op))
	log.Info("creating post", slog.String("userID", post.UserID.String()))

	if len(post.PostText) > 2500 {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectContents)
	}

	if len(post.Subject) > 100 {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectSubject)
	}

	if post.Price < 0 {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectPrice)
	}

	if !linkValidation(post.ImageAddress) {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectImageAddress)
	}

	resp, err := http.Head(post.ImageAddress)
	fmt.Println(resp)
	if err != nil || resp.Status != "200 OK" {

		return 0, fmt.Errorf("%s: %w", op, ErrHeadRequestFailed)
	}
	if resp.ContentLength > int64(5*1024*1024) {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectImageSize)
	}
	defer resp.Body.Close()

	id, err := u.postRepo.CreatePost(ctx, post)
	if err != nil {
		log.Error("failed to create post", slog.String("error", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("post created", slog.Int("postId", id))
	return id, nil
}

func (u *UserService) Posts(ctx context.Context, page, pageSize int, userId, sortBy, sortDir string, minPrice, maxPrice float64) (*[]dto.ResponsePost, error) {
	const op = "posts.service.Posts"
	log := u.log.With(slog.String("op", op))
	log.Info("getting posts")

	posts, err := u.postRepo.Posts(ctx, page, pageSize, userId, sortBy, sortDir, minPrice, maxPrice)
	if err != nil {
		log.Error("failed to get posts", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return posts, nil
}

func linkValidation(imageAddress string) bool {
	splitLink := strings.Split(imageAddress, ".")
	format := splitLink[len(splitLink)-1]
	_, ok := allowedPhotoTypes[format]
	return ok
}
