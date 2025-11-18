package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.back/configs"
	"go.back/internal/dto"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) CreateUser(request dto.Register) (string, error) {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (id, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", configs.UsersTable)

	row := a.db.QueryRow(query, uuid.New().String(), request.Username, request.Email, request.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
