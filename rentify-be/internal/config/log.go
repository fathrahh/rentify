package config

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	logFile, err := os.OpenFile("backend.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}

	multi := io.MultiWriter(os.Stdout, logFile)
	logger := zerolog.New(multi).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	return logger
}
