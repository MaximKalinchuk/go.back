package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.back/internal/dto"
	"go.back/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	signInKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL  = 24 * time.Hour
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
	passwordHash, err := s.GeneratePasswordHash(request.Password)

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
		return "", err
	}

	type tokenClaims struct {
		jwt.StandardClaims
		UserId uuid.UUID `json:"user_id"`
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) GeneratePasswordHash(password string) (string, error) {
	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(byteHashPassword), nil
}
