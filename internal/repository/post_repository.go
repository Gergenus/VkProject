package repository

import (
	"context"
	"fmt"

	"github.com/Gergenus/VkProject/internal/models"
)

func (p *PostgresRepository) CreatePost(ctx context.Context, post models.Post) (int, error) {
	const op = "post.repository.CreatePost"
	var id int

	err := p.db.DB.QueryRow(ctx, "INSERT INTO posts (user_id, subject, post_text, image_address, price) VALUES($1, $2, $3, $4, $5) RETURNING id",
		post.UserID.String(), post.Subject, post.PostText, post.ImageAddress, post.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
