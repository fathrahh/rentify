package http

import (
	"github.com/rs/zerolog"
	"ijor.dev/rentify/internal/usecase"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
	Log     zerolog.Logger
}

func NewUserHandler(userUsecase *usecase.UserUsecase, log zerolog.Logger) *UserHandler {
	return &UserHandler{
		Usecase: userUsecase,
		Log:     log,
	}
}
