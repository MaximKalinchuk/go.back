package service

import (
	"go.back/internal/dto"
	"go.back/internal/repository"
)

type Authorization interface {
	CreateUser(request dto.Register) (string, error)
	GenerateToken(request dto.Login) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.User),
	}
}
