package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"ijor.dev/rentify/internal/domain/entity"
	"ijor.dev/rentify/internal/domain/model"
	"ijor.dev/rentify/internal/domain/model/converter"
	"ijor.dev/rentify/internal/domain/repository"
)

type UserUsecase struct {
	UserRepo *repository.UserRepository
	Log      zerolog.Logger
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserUsecase(userRepo *repository.UserRepository, db *gorm.DB, log zerolog.Logger, validate *validator.Validate) *UserUsecase {
	log = log.With().Str("ctx", "user_usecase").Logger()
	log.Info().Msg("initialize user usecase")

	return &UserUsecase{
		UserRepo: userRepo,
		Log:      log,
		DB:       db,
		Validate: validate,
	}
}

func (c *UserUsecase) UserSignUp(ctx context.Context, payload *model.UserSignUpRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(payload)

	if err != nil {
		c.Log.Warn().Msgf("invalid request body : %+v", err)
	}

	total, err := c.UserRepo.CountByEmail(tx, payload.Email)

	if err != nil {
		c.Log.Warn().Msgf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warn().Msgf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		c.Log.Warn().Msgf("Failed to generate bcrypt hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		ID:    uuid.NewString(),
		Name:  payload.Name,
		Email: payload.Email,
		Hash:  string(hash),
	}

	if err := c.UserRepo.Create(tx, user); err != nil {
		c.Log.Warn().Msgf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warn().Msgf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	c.Log.Info().Msg("user created")

	return converter.UserEntityToResponse(user), nil
}
