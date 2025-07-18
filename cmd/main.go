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
	repo := repository.NewPostgresRepository(database)
	srv := service.NewUserService(repo, log, cfg.TokenTTL, cfg.JWTSecret)
	hnd := handlers.NewUserHandler(srv)

	e := echo.New()
	e.Use(middleware.Recover())

	auth := e.Group("/auth")
	{
		auth.POST("/signUp", hnd.SignUp)
		auth.POST("/signIn", hnd.SignIn)
	}

	e.GET("/", ZV, middlew.TestMiddleware)

	e.Start(":" + cfg.HTTPPort)

}

func ZV(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"uid":   c.Get("uid"),
		"login": c.Get("login"),
	})
}
