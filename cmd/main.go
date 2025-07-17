package main

import (
	"github.com/Gergenus/VkProject/internal/config"
	"github.com/Gergenus/VkProject/internal/repository"
	"github.com/Gergenus/VkProject/pkg/db"
	"github.com/Gergenus/VkProject/pkg/logger"
)

func main() {
	cfg := config.InitConfig()
	log := logger.SetupLogger(cfg.LogLevel, cfg.LogType)
	database := db.InitDB(cfg.PostgresURL)
	repo := repository.NewPostgresRepository(database)

}
