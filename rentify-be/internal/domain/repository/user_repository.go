package repository

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"ijor.dev/rentify/internal/domain/entity"
)

type UserRepository struct {
	Repository[entity.User]
	Log zerolog.Logger
}

func NewUserRepository(log zerolog.Logger) *UserRepository {
	return &UserRepository{
		Log: log.With().Str("ctx", "user_repository").Logger(),
	}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, entity *entity.User, email string) error {
	return db.Where("email = ?", email).Take(entity).Error
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Error
	return total, err

}
