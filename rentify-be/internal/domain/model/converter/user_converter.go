package converter

import (
	"ijor.dev/rentify/internal/domain/entity"
	"ijor.dev/rentify/internal/domain/model"
)

func UserEntityToResponse(entity *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
