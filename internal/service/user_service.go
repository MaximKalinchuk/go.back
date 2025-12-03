package service

import (
	"go.back/internal/entity"
	"go.back/internal/repository"
	"go.back/pkg/customerror"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id string) (entity.User, error) {
	user, err := s.repository.GetUserById(id)

	if err != nil {
		return entity.User{}, customerror.UserNotFound
	}

	return user, nil

}
