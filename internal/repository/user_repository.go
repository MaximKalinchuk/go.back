package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.back/configs"
	"go.back/internal/dto"
	"go.back/internal/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(request dto.Register) (string, error) {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (id, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", configs.UsersTable)

	row := u.db.QueryRow(query, uuid.New().String(), request.Username, request.Email, request.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserRepository) GetUser(email string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", configs.UsersTable)
	err := u.db.Get(&user, query, email)

	return user, err
}
