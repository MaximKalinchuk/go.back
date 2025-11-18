package service

import (
	"go.back/internal/dto"
	"go.back/internal/repository"
)

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) CreateUser(request dto.Register) (string, error) {
	userId, err := s.repository.CreateUser(request)
	return userId, err
}
