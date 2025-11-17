package service

import (
	"go.back/internal/dto"
	"go.back/internal/repository"
)

type AuthService struct {
	repository repository.Repository
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(request dto.Register) (int, error) {
	s.repository.CreateUser(request)
	return 0, nil
}
