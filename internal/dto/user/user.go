package userdto

import (
	"github.com/google/uuid"
	"go.back/internal/entity"
)

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func FromEntity(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
