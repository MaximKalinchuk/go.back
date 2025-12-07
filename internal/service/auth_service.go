package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	authdto "go.back/internal/dto/auth"
	"go.back/internal/repository"
	customerror "go.back/pkg/customerror"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"userId"`
}
type AuthService struct {
	repository repository.User
}

func NewAuthService(repository repository.User) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) CreateUser(request authdto.Register) (string, error) {
	passwordHash, err := s.generatePasswordHash(request.Password)

	if err != nil {
		return "", err
	}

	request.Password = passwordHash

	userId, err := s.repository.CreateUser(request)
	return userId, err
}

func (s *AuthService) GenerateToken(request authdto.Login) (string, error) {

	user, err := s.repository.GetUserByEmail(request.Email)

	if err != nil {
		return "", customerror.UserNotFound
	}

	err = s.checkPassword(request.Password, user.PasswordHash)

	if err != nil {
		return "", customerror.InvalidCredentials
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

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId.String(), nil
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
