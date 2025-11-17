package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.back/internal/dto"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) CreateUser(request dto.Register) (int, error) {
	fmt.Println("Данные в репозитории", request)
	return 1, nil
}
