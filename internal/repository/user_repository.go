package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type PostgresRepository struct {
	db db.PostgresDB
}

func NewPostgresRepository(db db.PostgresDB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) SaveUser(ctx context.Context, login string, passwordHash string) (models.User, error) {
	const op = "repository.SignUp"
	userID := uuid.New()
	_, err := p.db.DB.Exec(ctx, "INSERT INTO users (id, login, password_hash) VALUES($1, $2, $3)", userID.String(), login, passwordHash)
	if err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == "23505" {
				return models.User{}, fmt.Errorf("%s: %w", op, ErrUserExists)
			}
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return models.User{
		ID:    userID,
		Login: login,
	}, nil
}

func (p *PostgresRepository) User(ctx context.Context, login string) (models.User, error) {
	const op = "repository.User"
	var user models.User
	var stringUserID string
	err := p.db.DB.QueryRow(ctx, "SELECT * FROM users WHERE login = $1", login).Scan(&stringUserID, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	user.ID = uuid.MustParse(stringUserID)
	return user, nil
}
