package repository

import (
	"context"
	"fmt"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/transport/http/dto"
)

func (p *PostgresRepository) CreatePost(ctx context.Context, post models.ProductPost) (int, error) {
	const op = "post.repository.CreatePost"
	var id int

	err := p.db.DB.QueryRow(ctx, "INSERT INTO posts (user_id, subject, post_text, image_address, price) VALUES($1, $2, $3, $4, $5) RETURNING id",
		post.UserID.String(), post.Subject, post.PostText, post.ImageAddress, post.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (p *PostgresRepository) Posts(ctx context.Context, page, pageSize int, userId, sortBy, sortDir string, minPrice, maxPrice float64) (*[]dto.ResponsePost, error) {
	const op = "post.repository.Posts"
	var posts []dto.ResponsePost
	query := `
        SELECT 
            p.id, 
            p.user_id, 
            u.login,
            p.subject, 
            p.post_text, 
            p.image_address, 
            p.price, 
            p.created_at 
        FROM 
            posts p 
        JOIN 
            users u ON p.user_id = u.id
        WHERE 
            (COALESCE($1, 0) = 0 OR p.price >= $1) 
            AND (COALESCE($2, 0) = 0 OR p.price <= $2)
        ORDER BY
            CASE WHEN $3 = 'price' AND $4 = 'asc' THEN p.price END ASC,
            CASE WHEN $3 = 'price' AND $4 = 'desc' THEN p.price END DESC,
            CASE WHEN $3 = 'created_at' AND $4 = 'asc' THEN p.created_at END ASC,
            CASE WHEN $3 = 'created_at' AND $4 = 'desc' THEN p.created_at END DESC
        LIMIT $5 OFFSET ($6 - 1) * $5
    `
	rows, err := p.db.DB.Query(ctx, query, minPrice, maxPrice, sortBy, sortDir, pageSize, page)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for rows.Next() {
		var post dto.ResponsePost
		var uid string
		err = rows.Scan(&post.ID, &uid, &post.Login, &post.Subject, &post.PostText, &post.ImageAddress, &post.Price, &post.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		if uid == userId {
			post.IsOwner = true
		}
		posts = append(posts, post)
	}
	rows.Close()
	if rows.Err() != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &posts, nil
}
