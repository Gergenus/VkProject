package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(level, handlerType string) *slog.Logger {
	if handlerType == "json" {
		switch level {
		case "debug":
			return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case "prod":
			return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		default:
			panic("incorrect logger level")
		}
	}
	if handlerType == "text" {
		switch level {
		case "debug":
			return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case "prod":
			return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		default:
			panic("incorrect logger level")
		}
	}
	panic("incorrect handler type")
}
