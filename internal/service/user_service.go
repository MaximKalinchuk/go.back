package service

import (
	authdto "go.back/internal/dto/auth"
	"go.back/internal/entity"
	"go.back/internal/repository"
	"go.back/pkg/customerror"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(request authdto.Register) (string, error) {

	passwordHash, err := s.generatePasswordHash(request.Password)

	if err != nil {
		return "", nil
	}

	request.Password = passwordHash

	userId, err := s.repository.CreateUser(request)
	return userId, err
}

func (s *UserService) GetUser(id string) (entity.User, error) {
	user, err := s.repository.GetUserById(id)

	if err != nil {
		return entity.User{}, customerror.UserNotFound
	}

	return user, nil

}

func (s *UserService) generatePasswordHash(password string) (string, error) {
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(byteHashPassword), nil
}
