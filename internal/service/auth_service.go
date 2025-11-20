package service

import (
	"fmt"

	"go.back/internal/dto"
	"go.back/internal/repository"
)

type AuthService struct {
	repository repository.User
}

func NewAuthService(repository repository.User) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) CreateUser(request dto.Register) (string, error) {
	userId, err := s.repository.CreateUser(request)
	return userId, err
}

func (s *AuthService) GenerateToken(request dto.Login) (string, error) {

	user, err := s.repository.GetUser(request.Email)

	fmt.Print("\n", user, err)

	return "", nil
}
