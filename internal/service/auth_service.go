package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.back/internal/dto"
	"go.back/internal/repository"
	custom_error "go.back/pkg/customerror"
	"golang.org/x/crypto/bcrypt"
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
	passwordHash, err := s.generatePasswordHash(request.Password)

	if err != nil {
		return "", err
	}

	request.Password = passwordHash

	userId, err := s.repository.CreateUser(request)
	return userId, err
}

func (s *AuthService) GenerateToken(request dto.Login) (string, error) {

	user, err := s.repository.GetUser(request.Email)

	if err != nil {
		return "", custom_error.UserNotFound
	}

	err = s.checkPassword(request.Password, user.PasswordHash)

	if err != nil {
		return "", custom_error.InvalidCredentials
	}

	type tokenClaims struct {
		jwt.StandardClaims
		UserId uuid.UUID `json:"userId"`
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("jwt.token_ttl")).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(byteHashPassword), nil
}

func (s *AuthService) checkPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
