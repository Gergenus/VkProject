package main

import (
	"github.com/Gergenus/VkProject/internal/config"
	"github.com/Gergenus/VkProject/internal/handlers"
	"github.com/Gergenus/VkProject/internal/middlew"
	"github.com/Gergenus/VkProject/internal/repository"
	"github.com/Gergenus/VkProject/internal/service"
	"github.com/Gergenus/VkProject/pkg/db"
	"github.com/Gergenus/VkProject/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	log := logger.SetupLogger(cfg.LogLevel, cfg.LogType)
	database := db.InitDB(cfg.PostgresURL)
	defer database.DB.Close()
	repo := repository.NewPostgresRepository(database)
	srv := service.NewUserService(repo, repo, log, cfg.TokenTTL, cfg.JWTSecret)
	hnd := handlers.NewUserHandler(srv, srv)

	e := echo.New()
	e.Use(middleware.Recover())

	auth := e.Group("/auth")
	{
		auth.POST("/signUp", hnd.SignUp)
		auth.POST("/signIn", hnd.SignIn)
	}

	e.POST("/post/create", hnd.CreatePost, middlew.AuthMiddleware)
	e.GET("/posts", hnd.Posts, middlew.NotCompulsoryAuth)

	e.Start(":" + cfg.HTTPPort)
}
