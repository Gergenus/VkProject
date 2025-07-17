package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL string
	LogLevel    string
	HTTPPort    string
	JWTSecret   string
	LogType     string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		HTTPPort:    os.Getenv("HTTP_PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		LogType:     os.Getenv("LOG_TYPE"),
	}
}
