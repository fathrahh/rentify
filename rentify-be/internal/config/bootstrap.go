package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"ijor.dev/rentify/internal/delivery/http"
	"ijor.dev/rentify/internal/delivery/http/route"
	"ijor.dev/rentify/internal/domain/repository"
	"ijor.dev/rentify/internal/usecase"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	Log zerolog.Logger
	App *fiber.App
}

func Bootstrap(config BootstrapConfig) {
	// Repository
	userRepo := repository.NewUserRepository(config.Log)

	// Usecase
	userUc := usecase.NewUserUsecase(userRepo, config.DB, config.Log)

	// Handler
	userHandler := http.NewUserHandler(userUc, config.Log)

	// route config
	routeConfig := route.RouteConfig{
		App:         config.App,
		UserHandler: userHandler,
	}

	routeConfig.Setup()
}
