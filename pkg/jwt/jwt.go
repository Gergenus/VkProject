package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func GenerateNewToken(uid uuid.UUID, login string, duration time.Duration, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  uid.String(),
		"login": login,
		"exp":   time.Now().Add(duration).Unix(),
	})
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Returns uuid, login and error
func ParseToken(token, secret string) (string, string, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (any, error) { return []byte(secret), nil })
	if err != nil {
		return "", "", err
	}
	if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
		return "", "", ErrInvalidToken
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", ErrInvalidToken
	}

	uid, ok := claims["uuid"].(string)
	if !ok {
		return "", "", ErrInvalidToken
	}

	login, ok := claims["login"].(string)
	if !ok {
		return "", "", ErrInvalidToken
	}
	return uid, login, nil
}
