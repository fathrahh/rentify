package main

import (
	"github.com/joho/godotenv"
	"ijor.dev/rentify/internal/config"
)

func main() {
	godotenv.Load()
	logger := config.NewLogger()
	db, err := config.NewDB()

	if err != nil {
		logger.Error().Msg("failed to connect database")
	}

	app := config.NewFiberApp()
	config.Bootstrap(config.BootstrapConfig{
		DB: db,
	})

	logger.Info().Msg("server running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		logger.Error().Err(err)
	}
}
