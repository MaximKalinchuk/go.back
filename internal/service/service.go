package service

import (
	authdto "go.back/internal/dto/auth"
	"go.back/internal/entity"
	"go.back/internal/repository"
)

type Authorization interface {
	CreateUser(request authdto.Register) (string, error)
	GenerateToken(request authdto.Login) (string, error)
}

type User interface {
	GetUser(id string) (entity.User, error)
}

type Service struct {
	Authorization
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.User),
		User:          NewUserService(repository.User),
	}
}
