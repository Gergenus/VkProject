package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/repository"
	"github.com/Gergenus/VkProject/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      repository.UserRepositoryInterface
	log       *slog.Logger
	tokenTTL  time.Duration
	jwtSecret string
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user exists")
)

func NewUserService(repo repository.UserRepositoryInterface, log *slog.Logger, tokenTTL time.Duration, jwtSecret string) *UserService {
	return &UserService{
		repo:      repo,
		log:       log,
		tokenTTL:  tokenTTL,
		jwtSecret: jwtSecret,
	}
}

func (u *UserService) RegisterNewUser(ctx context.Context, login, password string) (models.User, error) {
	const op = "service.RegisterNewUser"

	log := u.log.With(
		slog.String("op", op),
		slog.String("login", login),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err := u.repo.SaveUser(ctx, login, string(passHash))
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			log.Warn("user exists", slog.String("login", login))
			return models.User{}, fmt.Errorf("%s: %w", op, err)
		}
		log.Error("failed to save user", slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("user has been created", slog.String("login", login))
	return user, nil
}

func (u *UserService) Login(ctx context.Context, login, password string) (string, error) {
	const op = "service.Login"
	log := u.log.With(slog.String("op", op), slog.String("login", login))

	log.Info("logging the user")

	user, err := u.repo.User(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Warn("user not found", slog.String("login", login), slog.String("error", err.Error()))
			return "", fmt.Errorf("%s; %w", op, ErrInvalidCredentials)
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Info("invalid credentials", slog.String("login", login), slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully", slog.String("login", login))

	token, err := jwt.GenerateNewToken(user.ID, user.Login, u.tokenTTL, u.jwtSecret)
	if err != nil {
		log.Error("failed to generate token", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}
